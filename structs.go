package main

import "encoding/xml"

type Devices struct {
	XMLName xml.Name `xml:"devices"`
	Devices []Device `xml:"device"`
}

type Device struct {
	DeviceID string `xml:"device,attr"`
	Type     string `xml:"type,attr"`
	Function string `xml:"function,attr"`
	Bus      string `xml:"bus,attr"`
	Config   Config `xml:"config_space"`
}

type Config struct {
	Bytes string `xml:"bytes"`
}

type PCIeHeader struct {
	VendorID                uint16
	DeviceID                uint16
	Command                 uint16
	Status                  uint16
	RevisionID              uint8
	ClassCode               [3]byte
	CacheLineSize           uint8
	LatencyTimer            uint8
	HeaderType              uint8
	BuiltInSelfTest         uint8
	BaseRegister0           uint32
	BaseRegister1           uint32
	BaseRegister2           uint32
	BaseRegister3           uint32
	BaseRegister4           uint32
	BaseRegister5           uint32
	CardbusCISPointer       uint32
	SubsystemVendorID       uint16
	SubsystemID             uint16
	ExpansionROMBaseAddress uint32
	CapabilitiesPointer     uint8
	Reserved1               [3]byte
	Reserved2               uint32
	InterruptLine           uint8
	InterruptPin            uint8
	MinGrant                uint8
	MaxLatency              uint8
}

type CapabilitiesInfo struct {
	Name      string
	Size      int
	WriteMask []uint32
}

// PCI Code and ID Assignment Specification Revision 1.11
var CAPABILITIES map[uint8]CapabilitiesInfo = map[uint8]CapabilitiesInfo{
	0x00: {"NULL", 0, nil},
	0x01: {"PCI Power Management Interface ", len(Write_protected_bits_PM), Write_protected_bits_PM},
	0x02: {"Accelerated Graphics Port (DEPRECATED)", 0, nil},
	0x03: {"Vital Product Data", len(Write_protected_bits_VPD), Write_protected_bits_VPD},
	0x04: {"Slot Identification", 8, nil},
	0x05: {"Message Signaled Interrupts", -1, nil},
	0x06: {"CompactPCI Hot Swap", 8, nil},
	0x07: {"PCI-X", 0, nil},
	0x08: {"HyperTransport ", 0, nil},
	0x09: {"Vendor Specific", len(Write_protected_bits_VSC), Write_protected_bits_VSC},
	0x0A: {"Debug port", 16, nil},
	0x0B: {"CompactPCI central resource control ", 0, nil},
	0x0C: {"PCI Hot-Plug", 8, nil},
	0x0D: {"PCI Bridge Subsystem Vendor ID", 8, nil},
	0x0E: {"AGP 8x", 0, nil},
	0x0F: {"Secure Device", 0, nil},
	0x10: {"PCI Express", len(Write_protected_bits_PCIE), Write_protected_bits_PCIE},
	0x11: {"MSI-X", -2, nil},
	0x12: {"Serial ATA Data/Index Configuration", 8, nil},
	0x13: {"Advanced Features", 0, nil},
	0x14: {"Enhanced Allocation", 0, nil},
	0x15: {"Flattening Portal Bridge", 0, nil},
}

// PCI Code and ID Assignment Specification Revision 1.11
var EXTENDED_CAPABILITIES map[int]CapabilitiesInfo = map[int]CapabilitiesInfo{
	0x0001: {"advanced error reporting", len(Write_protected_bits_AER), Write_protected_bits_AER},
	0x0002: {"virtual channel", 0, nil},
	0x0003: {"device serial number", len(Write_protected_bits_DSN), Write_protected_bits_DSN},
	0x0004: {"power budgeting", 2, nil},
	0x0005: {"root complex link declaration", 0, nil},
	0x0006: {"root complex internal link control", 3, nil},
	0x0007: {"root complex event collector endpoint association", 3, nil},
	0x0008: {"multi-function virtual channel", 0, nil},
	0x0009: {"virtual channel", 0, nil},
	0x000A: {"root complex register block", len(Write_protected_bits_VSEC), Write_protected_bits_VSEC},
	0x000B: {"vendor specific", len(Write_protected_bits_PTM), Write_protected_bits_PTM},
	0x000C: {"configuration access correlation", 3, nil},
	0x000D: {"access control services", 0, nil},
	0x000E: {"alternative routing-ID interpretation", 2, nil},
	0x000F: {"address translation services", 0, nil},
	0x0010: {"single root IO virtualization", 0, nil},
	0x0011: {"multi-root IO virtualization", 0, nil},
	0x0012: {"multicast", 0, nil},
	0x0013: {"page request interface", 2, nil},
	0x0014: {"AMD reserved", 0, nil},
	0x0015: {"resizable BAR", 0, nil},
	0x0016: {"dynamic power allocation", 4, nil},
	0x0017: {"TPH requester", len(Write_protected_bits_TPH), Write_protected_bits_TPH},
	0x0018: {"latency tolerance reporting", 2, nil},
	0x0019: {"secondary PCI express", 0, nil},
	0x001A: {"protocol multiplexing", 3, nil},
	0x001B: {"process address space ID", 2, nil},
	0x001C: {"LN requester", 2, nil},
	0x001D: {"downstream port containment", 2, nil},
	0x001E: {"L1 PM substates", len(Write_protected_bits_L1PM), Write_protected_bits_L1PM},
	0x001F: {"precision time measurement", 2, nil},
	0x0020: {"M-PCIe", 0, nil},
	0x0021: {"FRS queueing", 3, nil},
	0x0022: {"Readiness time reporting", 2, nil},
	0x0023: {"designated vendor specific", 0, nil},
	0x0024: {"VF resizable BAR", 0, nil},
	0x0025: {"data link feature", 3, nil},
	0x0026: {"physical layer 16.0 GT/s", 4, nil},
	0x0027: {"receiver lane margining", 0, nil},
	0x0028: {"hierarchy ID", 2, nil},
	0x0029: {"native PCIe enclosure management", 0, nil},
	0x002A: {"physical layer 32.0 GT/s", 3, nil},
	0x002B: {"alternate protocol", 0, nil},
	0x002C: {"system firmware intermediary", 0, nil},
}

var PCIeHeaderReadWriteMask [16]uint32 = [16]uint32{
	0x00000000, 0x470500f9, 0x00000000, 0xffff0040,
	0xf0ffffff, 0xffffffff, 0xf0ffffff, 0xffffffff,
	0xf0ffffff, 0xf0ffffff, 0x00000000, 0x00000000,
	0x01f8ffff, 0x00000000, 0x00000000, 0xff000000}
