package libdns

func readBoolFromUint16(row uint16, shift int) bool {
	PrintlnUint16WithZeroPad(row&(0b1<<shift)>>shift)
	println(row&(0b1<<shift)>>shift==0b1)
	return row&(0b1<<shift)>>shift == 0b1
}

func readQRBit(row uint16) bool {
	return readBoolFromUint16(row, 15)

}

func readOpcode(row uint16) uint8 {
	return uint8(row & (0b1111 << 11) >> 11)
}

func readAABit(row uint16) bool {
	return readBoolFromUint16(row, 10)
}
func readTCBit(row uint16) bool {
	return readBoolFromUint16(row, 9)
}

func readRDBit(row uint16) bool {
	return readBoolFromUint16(row, 8)
}

func readRABit(row uint16) bool {
	return readBoolFromUint16(row, 7)
}

func readZBit(row uint16) bool {
	return readBoolFromUint16(row, 6)
}

func readADBit(row uint16) bool {
	return readBoolFromUint16(row, 5)
}

func readCDBit(row uint16) bool {
	return readBoolFromUint16(row, 4)
}

func readRCode(row uint16) uint8 {
	return uint8(row & 0b1111)
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
func i2b(b int) bool {
	return 1 == b
}
