package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	source := strings.NewReader("dadadadadasd")
	buffered := bufio.NewReader(source)
	newstring, err := buffered.ReadString('\n')
	if err == io.EOF {
		fmt.Println(newstring)
	} else {
		fmt.Println("something went wrong...")
	}
}
