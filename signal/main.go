package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch := make(chan os.Signal)
	quit := make(chan struct{})
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, os.Interrupt)
	go func() {
		<-ch
		quit <- struct{}{}
	}()
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Good")
	case <-quit:
		fmt.Println("interrupt")
	}
}
