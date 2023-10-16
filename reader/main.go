package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("sdasdasdasd")
	var newString strings.Builder
	buffer := make([]byte, 4)
	for {
		numBytes, err := reader.Read(buffer)
		chunk := buffer[:numBytes]
		newString.Write(chunk)
		fmt.Printf("Read %v bytes: %c\n", numBytes, chunk)
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("%v\n", newString.String())
}
