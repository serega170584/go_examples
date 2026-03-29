package main

import (
	"fmt"
	"time"
)

func main() {
	var a []int

	go func() {
		defer func() {
			fmt.Println("1234")
		}()
		a[1] = 1
	}()

	time.Sleep(time.Second)
}
