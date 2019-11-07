package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	var arr []uint8
	var i uint32

	// ---- big endian ↓ ----
	arr = []uint8{0, 0, 0, 1}
	i = binary.BigEndian.Uint32(arr)
	fmt.Printf("big-endian: %d\n", i)

	arr = []uint8{0, 0, 1, 1}
	i = binary.BigEndian.Uint32(arr)
	fmt.Printf("big-endian: %d\n", i)

	// ---- little endian ↓ ----
	arr = []uint8{1, 0, 0, 0}
	i = binary.LittleEndian.Uint32(arr)
	fmt.Printf("little-endian: %d\n", i)

	arr = []uint8{1, 1, 0, 0}
	i = binary.LittleEndian.Uint32(arr)
	fmt.Printf("little-endian: %d\n", i)
}
