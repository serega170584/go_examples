package main

import (
	"fmt"
	"time"
)

func main() {
	select {
	case <-time.After(time.Nanosecond):
		fmt.Println("Second")
	case <-time.After(2 * time.Nanosecond):
		fmt.Println("3 seconds")
	}
}
