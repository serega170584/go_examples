package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	t := time.NewTicker(2 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	func() {
		for {
			select {
			case <-t.C:
				fmt.Println("Result " + strconv.Itoa(rand.Intn(1000)))
			case <-ctx.Done():
				t.Stop()
				return
			}
		}
	}()
}
