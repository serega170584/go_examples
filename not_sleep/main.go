package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx := context.Background()
	err := Run(ctx)
	if err != nil {
		fmt.Println(err)
	}
}

func Run(ctx context.Context) error {
	for {
		a := rand.Intn(100)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(3 * time.Second):
			b := rand.Intn(1000)
		}
	}
	return nil
}
