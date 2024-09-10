# PCIeConfigSpace

This utility interprets PCIe configuration space data extracted by [TeleScan PE Software](https://www.teledynelecroy.com/protocolanalyzer/pci-express/telescan-pe-software). It serves two main functions:

1. Analyzing the configuration space
2. Generating a corresponding write mask

The tool produces two output files compatible with pcileech-fpga:

1. A shadow configuration space
2. A write mask

These files can be used to enhance the legitimacy of your DMA card's appearance when utilizing [pcileech-fpga](https://github.com/ufrisk/pcileech-fpga).

## Usage

```shell
PS C:\Tools\PCIeConfigSpace> .\TLScan.exe -h
Usage of C:\Tools\PCIeConfigSpace\TLScan.exe:
  -i string
        The .tlscan scan file of your donor card (default "donor.tlscan")
  -overwrite
        Automatically overwrite existing file
  -so string
        The output .coe file containing your formatted config space (default "pcileech_cfgspace_extracted.coe")
  -wo string
        The output .coe file containing your formatted config space writemask (default "pcileech_cfgspace_writemask_extracted.coe")
```

### Resources

- PCI Local Bus Specification (Revision 2.2)
- PCI Code and ID Assignment Specification (Revision 1.11)
- PCI Express® Base Specification (Revision 3.0)
- [Writemask.it by Simonrak](https://github.com/Simonrak/writemask.it)
