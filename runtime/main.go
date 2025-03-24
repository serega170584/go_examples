package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				i++
			}
		}()
	}
	fmt.Println(runtime.LockOSThread())
}
