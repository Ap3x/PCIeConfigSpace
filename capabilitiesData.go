package main

import (
	"bytes"
	"encoding/binary"
)

func GetCapabilityWriteMask(capabilityID uint8) []byte {
	capInfo := CAPABILITIES[capabilityID]
	return Uint32SliceToBytes(capInfo.WriteMask)
}

func GetCapabilityID(capabilityBase []byte) byte {
	return capabilityBase[0]
}

func GetCapabilityNextPointer(capabilityBase []byte) byte {
	return capabilityBase[1]
}

func GetCapabilityName(capabilityID uint8) string {
	capInfo := CAPABILITIES[capabilityID]
	return capInfo.Name
}

func GetCapabilitySize(capabilityBase []byte) int {
	capabilityID := GetCapabilityID(capabilityBase)
	capInfo := CAPABILITIES[capabilityID]

	if capInfo.Size == -1 { //MSI 14bytes or 24 bytes
		var messageControl uint16
		err := binary.Read(bytes.NewReader(capabilityBase[2:]), binary.LittleEndian, &messageControl)
		if err != nil {
			return -1
		}
		bitArray := Uint16ToBitArray(messageControl)
		if !bitArray[8] { // If 64 bit disabled
			return 1
		} else if bitArray[8] && bitArray[7] { // If 64 bit enabled and pre-vector masking is true
			return 6
		} else {
			return 4
		}

	} else if capInfo.Size == -2 { //MSI-X - get TableSize
		var messageControl uint16
		err := binary.Read(bytes.NewReader(capabilityBase[2:]), binary.LittleEndian, &messageControl)
		if err != nil {
			return -1
		}
		bitArray := Uint16ToBitArray(messageControl)
		bitArrayTableSize := bitArray[6:]
		tableSize := BinaryArrayToDecimal(bitArrayTableSize) + 1
		return tableSize / 2

	} else {
		return capInfo.Size
	}
}
