package libdns

import (
	"fmt"
)

func PrintlnUint8WithZeroPad(value uint8) {
	fmt.Printf("%08b\n", value)
}
func PrintlnUint16WithZeroPad(value uint16) {
	fmt.Printf("%016b\n", value)
}
func PrintlnUint32WithZeroPad(value uint32) {
	fmt.Printf("%032b\n", value)
}
func PrintlnUint64WithZeroPad(value uint64) {
	fmt.Printf("%064b\n", value)
}
