package main

import (
	"encoding/binary"
)

// Write Mask Protected Bits from https://github.com/Simonrak/writemask.it/blob/main/WritemaskerTM.py
var Write_protected_bits_PCIE []uint32 = []uint32{
	0x00000f00,
	0x00000010,
	0xff7f0f00,
	0x00000000,
	0xcb0d00c0,
	0x00000000,
	0x0000ffff,
	0x00000000,
	0x00000000,
	0x00000000,
	0xff7f0000,
	0x00000000,
	0xbfff2000,
}

var Write_protected_bits_PM []uint32 = []uint32{
	0x00000000,
	0x00000000,
}

var Write_protected_bits_MSI_ENABLED_0 []uint32 = []uint32{
	0x0000f104,
}

var Write_protected_bits_MSI_64_bit_1 []uint32 = []uint32{
	0x0000f104,
	0x03000000,
	0x00000000,
	0xffff0000,
}

var Write_protected_bits_MSI_Multiple_Message_Capable_1 []uint32 = []uint32{
	0x0000f104,
	0x03000000,
	0x00000000,
	0xffff0000,
	0x00000000,
	0x01000000,
}

var Write_protected_bits_MSIX_3 []uint32 = []uint32{
	0x000000c0,
	0x00000000,
	0x00000000,
}

var Write_protected_bits_MSIX_4 []uint32 = []uint32{
	0x000000c0,
	0x00000000,
	0x00000000,
	0x00000000,
}

var Write_protected_bits_MSIX_5 []uint32 = []uint32{
	0x000000c0,
	0x00000000,
	0x00000000,
	0x00000000,
	0x00000000,
}

var Write_protected_bits_MSIX_6 []uint32 = []uint32{
	0x000000c0,
	0x00000000,
	0x00000000,
	0x00000000,
	0x00000000,
	0x00000000,
}

var Write_protected_bits_MSIX_7 []uint32 = []uint32{
	0x000000c0,
	0x00000000,
	0x00000000,
	0x00000000,
	0x00000000,
	0x00000000,
	0x00000000,
}

var Write_protected_bits_MSIX_8 []uint32 = []uint32{
	0x000000c0,
	0x00000000,
	0x00000000,
	0x00000000,
	0x00000000,
	0x00000000,
	0x00000000,
	0x00000000,
}

var Write_protected_bits_VPD []uint32 = []uint32{
	0x0000ffff,
	0xffffffff,
}

var Write_protected_bits_VSC []uint32 = []uint32{
	0x000000ff,
	0xffffffff,
}

var Write_protected_bits_TPH []uint32 = []uint32{
	0x00000000,
	0x00000000,
	0x070c0000,
}

var Write_protected_bits_VSEC []uint32 = []uint32{
	0x00000000,
	0x00000000,
	0xffffffff,
	0xffffffff,
}

var Write_protected_bits_AER []uint32 = []uint32{
	0x00000000,
	0x31f0ff07,
	0x31f0ff07,
	0x31f0ff07,
	0xc1f10000,
	0xc1f10000,
	0x40050000,
	0x00000000,
	0x00000000,
	0x00000000,
	0x00000000,
}

var Write_protected_bits_DSN []uint32 = []uint32{
	0x00000000,
	0x00000000,
	0x00000000,
}

var Write_protected_bits_LTR []uint32 = []uint32{
	0x00000000,
	0x00000000,
}

var Write_protected_bits_L1PM []uint32 = []uint32{
	0x00000000,
	0x00000000,
	0x3f00ffe3,
	0xfb000000,
}

var Write_protected_bits_PTM []uint32 = []uint32{
	0x00000000,
	0x00000000,
	0x00000000,
	0x03ff0000,
}

func convertToBytes(arr [16]uint32) []byte {
	result := make([]byte, len(arr)*4) // Each uint32 is 4 bytes
	for i, _ := range arr {
		binary.LittleEndian.PutUint32(result[i*4:], arr[i])
	}
	return result
}

func SetPCIeHeaderWriteMask(writeMaskBytes *[]byte) {
	pcieWriteMask := convertToBytes(PCIeHeaderReadWriteMask)
	copy((*writeMaskBytes)[:], pcieWriteMask)

}

func SetWriteMaskOffset(writeMaskBytes *[]byte, offset int, data []byte) {
	copy((*writeMaskBytes)[offset:], data)
}
