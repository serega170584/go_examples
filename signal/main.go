package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sign := make(chan os.Signal)
	res := make(chan struct{})
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)
	go func() {
		<-sign
		fmt.Println("Close")
		res <- struct{}{}
	}()
	<-res
}
