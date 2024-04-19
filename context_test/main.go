package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Err struct{}

func (e Err) Error() string {
	return "dfsdfsdf"
}

func err() *Err {
	return nil
}

func main() {
	//var a *int
	//var b int64
	fmt.Println(err() == nil)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	err := adapter(ctx)
	if err != nil {
		fmt.Println(err)
	}
	cancel()
}

func test() {
	time.Sleep(time.Duration(rand.Intn(6)) * time.Second)
	fmt.Println("Success")
}

func adapter(ctx context.Context) error {
	res := make(chan struct{})
	go func() {
		test()
		close(res)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-res:
		return nil
	}
}
