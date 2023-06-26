package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type resp struct {
	err     error
	content string
}

func main() {
	res := make(chan resp)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go rpcCall(ctx, res)
	output := <-res
	fmt.Println(output)
}

func rpcCall(ctx context.Context, res chan resp) {
	select {
	case <-ctx.Done():
		res <- resp{err: errors.New("Error")}
	case <-time.After(time.Duration(rand.Intn(7)) * time.Second):
		res <- resp{content: "Good"}
	}
}
