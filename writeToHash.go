package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	h := crc32.NewIEEE()
	fmt.Fprintf(h, "Hello Go\n")
	fmt.Printf("hash=%d\n", h.Sum32())
	fmt.Printf("hash=%#x\n", h.Sum32())
}
