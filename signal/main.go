package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch := make(chan os.Signal)
	ctx, cancel := context.WithCancel(context.Background())
	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-ch
		fmt.Println("123123123")
		time.Sleep(3 * time.Second)
		cancel()
	}()
	<-ctx.Done()
}
