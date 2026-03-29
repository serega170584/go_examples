package main

import (
	"fmt"
	"time"
)

func main() {
	a, _ := time.Parse(time.RFC3339, "2025-12-11T15:00:00Z")
	loc := a.Location()

	target := time.Now()
	targetTrMinute := target.In(loc).Truncate(time.Minute)

	//t := info.DeliveryDate.Truncate(time.Minute)
	//t := time.Now()
	fmt.Println(a.After(targetTrMinute))
}
