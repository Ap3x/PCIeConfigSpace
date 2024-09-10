package main

import (
	"encoding/binary"
	"fmt"
	"strings"
	"unicode"
)

func RemoveWhitespace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func Uint32SliceToBytes(input []uint32) []byte {
	result := make([]byte, len(input)*4)
	for i, v := range input {
		binary.LittleEndian.PutUint32(result[i*4:], v)
	}
	return result
}

func Uint16ToBitArray(num uint16) [16]bool {
	var bitArray [16]bool
	for i := 0; i < 16; i++ {
		bitArray[15-i] = (num & (1 << i)) != 0
	}
	return bitArray
}

func byteToUint32LittleEndian(b []byte) []uint32 {
	// Ensure the byte slice length is a multiple of 4
	if len(b)%4 != 0 {
		panic("Byte slice length must be a multiple of 4")
	}

	result := make([]uint32, len(b)/4)

	for i := 0; i < len(b); i += 4 {
		result[i/4] = binary.LittleEndian.Uint32(b[i : i+4])
	}

	return result
}

func byteToUint32BigEndian(b []byte) []uint32 {
	// Ensure the byte slice length is a multiple of 4
	if len(b)%4 != 0 {
		panic("Byte slice length must be a multiple of 4")
	}

	result := make([]uint32, len(b)/4)

	for i := 0; i < len(b); i += 4 {
		result[i/4] = binary.BigEndian.Uint32(b[i : i+4])
	}

	return result
}

func BinaryArrayToDecimal(binaryArray []bool) int {
	result := 0
	for i, bit := range binaryArray {
		if bit {
			result += 1 << (len(binaryArray) - 1 - i)
		}
	}
	return result
}

func HexStringToBytes(hexStr string) ([]byte, error) {
	var bytes []byte
	length := len(hexStr)
	if length%2 != 0 {
		return nil, fmt.Errorf("invalid hex string length")
	}
	for i := 0; i < length; i += 2 {
		var b uint8
		_, err := fmt.Sscanf(hexStr[i:i+2], "%2x", &b)
		if err != nil {
			return nil, err
		}
		bytes = append(bytes, byte(b))
	}
	return bytes, nil
}

// Read3BytesLittleEndian reads 3 bytes from a slice in little-endian order
func Format3BytesLittleEndian(dataToFormat ...*[3]byte) error {

	for i, data := range dataToFormat {
		if len(data) < 3 {
			return fmt.Errorf("not enough data to format in %d: need 3 bytes, got %d", i, len(data))
		}
		var tempData [3]byte = [3]byte{data[2], data[1], data[0]}
		*data = tempData
	}

	return nil
}
