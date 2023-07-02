package main

import (
	"context"
	"fmt"
	"math/rand"
)

func main() {
	ctx := context.Background()
	resp, err := somethingAdapter(ctx)
	if err != nil {
		fmt.Println(err)
	}
}

func something() {
	for {
		a := rand.Intn(1000)
	}
}

func somethingAdapter(ctx context.Context) (int, error) {
	ch := make(chan struct{})

	go func() {
		something()
		close(ch)
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case <-ch:
		return rand.Intn(1000), nil
	}
}
