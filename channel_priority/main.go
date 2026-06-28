package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chList := make([]chan *int, 2)
	for i := 0; i < 2; i++ {
		chList[i] = make(chan *int, 1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	v, err := func(ctx context.Context) (*int, error) {
		go func() {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			v1 := rand.Intn(100)
			fmt.Println(fmt.Sprintf("channel 1 value: %d", v1))
			chList[0] <- &v1
		}()

		go func() {
			time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
			v2 := rand.Intn(100)
			fmt.Println(fmt.Sprintf("channel 2 value: %d", v2))
			chList[1] <- &v2
		}()

		select {
		case <-ctx.Done():
			fmt.Println("context error")
			for _, c := range chList {
				cv := <-c
				if cv != nil {
					return cv, nil
				}
			}
			return nil, ctx.Err()
		case cv := <-chList[0]:
			return cv, nil
		}
	}(ctx)

	if err != nil {
		fmt.Println(fmt.Errorf("getting value error: %w", err))
		return
	}

	fmt.Println(*v)
}
