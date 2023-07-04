package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, os.Interrupt)
	go func() {
		<-ch
		fmt.Println(rand.Intn(1000))
		cancel()
	}()
	<-ctx.Done()
}
