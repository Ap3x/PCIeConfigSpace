package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

func ReadPCIeHeader(data []byte, offset int) (PCIeHeader, error) {
	var header PCIeHeader
	size := int(unsafe.Sizeof(header))

	if len(data) < offset+size {
		return header, fmt.Errorf("not enough data to read PCIeHeader at offset %d", offset)
	}

	err := binary.Read(bytes.NewReader(data[offset:offset+size]), binary.LittleEndian, &header)
	if err != nil {
		return header, fmt.Errorf("failed to read PCIeHeader: %v", err)
	}

	err = Format3BytesLittleEndian(&header.Reserved1, &header.ClassCode)
	if err != nil {
		return header, fmt.Errorf("failed to format three byte array: %v", err)
	}

	return header, nil
}
