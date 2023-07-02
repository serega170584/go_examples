package main

import (
	"context"
	"fmt"
	"math/rand"
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
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		a := rand.Intn(100)
		return nil
	}
}
