package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse("2006-01-02 15:04:05Z07:00", "2027-12-01 00:00:00+03:00")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
}
