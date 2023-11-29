package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := predictableFunc(ctx)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Println(res)
}

func unpredictableFunc() int {
	res := rand.Intn(5)
	time.Sleep(time.Duration(res) * time.Second)
	return res
}

func predictableFunc(ctx context.Context) (int, error) {
	res := make(chan int)

	go func() {
		res <- unpredictableFunc()
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case <-res:
		return unpredictableFunc(), nil
	}
}
