package utils

func ReadBoolFromUint16(row uint16, shift int) bool {
	return row&(0b1<<shift)>>shift == 0b1
}

func ReadQRBit(row uint16) bool {
	return ReadBoolFromUint16(row, 15)

}

func ReadOpcode(row uint16) uint8 {
	return uint8(row & (0b1111 << 11) >> 11)
}

func ReadAABit(row uint16) bool {
	return ReadBoolFromUint16(row, 10)
}
func ReadTCBit(row uint16) bool {
	return ReadBoolFromUint16(row, 9)
}

func ReadRDBit(row uint16) bool {
	return ReadBoolFromUint16(row, 8)
}

func ReadRABit(row uint16) bool {
	return ReadBoolFromUint16(row, 7)
}

func ReadZBit(row uint16) bool {
	return ReadBoolFromUint16(row, 6)
}

func ReadADBit(row uint16) bool {
	return ReadBoolFromUint16(row, 5)
}

func ReadCDBit(row uint16) bool {
	return ReadBoolFromUint16(row, 4)
}

func ReadRCode(row uint16) uint8 {
	return uint8(row & 0b1111)
}

func B2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
func I2b(b int) bool {
	return 1 == b
}
