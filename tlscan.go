package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type CommandLineArgs struct {
	inputTlScanFile     *string
	shadowOutputFile    *string
	writeMaskOutputFile *string
	overwriteAuto       *bool
}

func createWriteMask(sizeInBytes int) []byte {
	var writeMask []byte
	for i := 0; i < sizeInBytes; i++ {
		writeMask = append(writeMask, 0xFF)
	}

	return writeMask
}

func parseCmdLine() CommandLineArgs {
	var results CommandLineArgs

	inputHelp := "The .tlscan scan file of your donor card"
	defaultInput := "donor.tlscan"
	results.inputTlScanFile = flag.String("i", defaultInput, inputHelp)

	shadowOutputHelp := "The output .coe file containing your formatted config space"
	defaultShadowOutput := "pcileech_cfgspace_extracted.coe"
	results.shadowOutputFile = flag.String("so", defaultShadowOutput, shadowOutputHelp)

	writemaskOutputHelp := "The output .coe file containing your formatted config space writemask"
	defaultWriteMaskOutput := "pcileech_cfgspace_writemask_extracted.coe"
	results.writeMaskOutputFile = flag.String("wo", defaultWriteMaskOutput, writemaskOutputHelp)

	overwriteHelp := "Automatically overwrite existing file"
	results.overwriteAuto = flag.Bool("overwrite", false, overwriteHelp)

	flag.Parse()
	return results
}

func main() {
	cmdArgs := parseCmdLine()

	tlscanFile, err := os.Open(*cmdArgs.inputTlScanFile)
	if err != nil {
		fmt.Println("Error opening .tlscan file:", err)
		return
	}

	defer tlscanFile.Close()

	content, err := io.ReadAll(tlscanFile)
	if err != nil {
		fmt.Println("Error reading contents of .tlscan file:", err)
		return
	}

	var devices Devices
	err = xml.Unmarshal(content, &devices)
	if err != nil {
		fmt.Println("Error unmarshaling XML:", err)
		return
	}

	for _, device := range devices.Devices {
		cfgSpace, err := HexStringToBytes(RemoveWhitespace(device.Config.Bytes))
		if err != nil {
			fmt.Println("Error converting hex string to bytes:", err)
			return
		}

		writeMaskBytes := createWriteMask(len(cfgSpace))

		SetPCIeHeaderWriteMask(&writeMaskBytes)

		header, err := ReadPCIeHeader(cfgSpace, 0)
		if err != nil {
			fmt.Printf("Error Reading PCIe Header: %v\n", err)
			return
		}

		fmt.Printf("Vendor ID: 0x%04X\n", header.VendorID)
		fmt.Printf("Device ID: 0x%04X\n", header.DeviceID)
		fmt.Printf("Subsystem Vendor ID: 0x%02X\n", header.SubsystemVendorID)
		fmt.Printf("Subsystem ID: 0x%02X\n", header.SubsystemID)
		fmt.Printf("Revision: 0x%02X\n", header.RevisionID)
		fmt.Printf("Class Code: 0x%X\n", header.ClassCode)
		fmt.Printf("BAR0: 0x%X\n", header.BaseRegister0)
		fmt.Printf("BAR1: 0x%X\n", header.BaseRegister1)
		fmt.Printf("BAR2: 0x%X\n", header.BaseRegister2)
		fmt.Printf("BAR3: 0x%X\n", header.BaseRegister3)
		fmt.Printf("BAR4: 0x%X\n", header.BaseRegister4)
		fmt.Printf("BAR5: 0x%X\n", header.BaseRegister5)

		capIndex := int(header.CapabilitiesPointer)
		capPtr := cfgSpace[capIndex:]
		fmt.Printf("Capabilities:\n")
		for {
			capabilityID := GetCapabilityID(capPtr)
			dwordSize := GetCapabilitySize(capPtr)
			capabilityName := GetCapabilityName(capabilityID)
			fmt.Printf("\t%s - %d bytes\n", capabilityName, dwordSize*4)

			capWriteMask := GetCapabilityWriteMask(capabilityID)

			SetWriteMaskOffset(&writeMaskBytes, capIndex, capWriteMask)

			capIndex = int(GetCapabilityNextPointer(capPtr))
			if capIndex == 0 {
				break
			}

			capPtr = cfgSpace[capIndex:]
		}

		cfgSpaceBytesUint32 := byteToUint32BigEndian(cfgSpace)
		writeFile(cfgSpaceBytesUint32, *cmdArgs.shadowOutputFile, *cmdArgs.overwriteAuto)

		writeMaskBytesUint32 := byteToUint32LittleEndian(writeMaskBytes)
		writeFile(writeMaskBytesUint32, *cmdArgs.writeMaskOutputFile, *cmdArgs.overwriteAuto)
	}
}

func writeFile(slice []uint32, outputFile string, overwrite bool) {

	// Check if file exists
	_, err := os.Stat(outputFile)
	if err == nil {
		if !overwrite {
			fmt.Printf("File %s already exists. Overwrite? (y/n): ", outputFile)
			reader := bufio.NewReader(os.Stdin)
			response, _ := reader.ReadString('\n')
			response = strings.TrimSpace(strings.ToLower(response))

			if response != "y" && response != "yes" {
				fmt.Println("Operation cancelled.")
				return
			}
		}
	} else if !os.IsNotExist(err) {
		fmt.Println("Error checking file:", err)
		return
	}

	var builder strings.Builder
	builder.WriteString("memory_initialization_radix=16;\nmemory_initialization_vector=\n\n")

	for i, value := range slice {
		builder.WriteString(fmt.Sprintf("%08x", value))

		if i == len(slice)-1 {
			builder.WriteString("\n")
		} else if (i+1)%4 == 0 {
			builder.WriteString("\n")
		} else {
			builder.WriteString(",")
		}
	}

	err = os.WriteFile(outputFile, []byte(builder.String()), 0644)
	if err != nil {
		fmt.Printf("Error writing to %s: %s\n", outputFile, err)
		return
	}

	fmt.Printf("Successfully wrote %s\n", outputFile)
}
