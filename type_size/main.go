package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s [1]int
	fmt.Println(unsafe.Sizeof(s))
}
