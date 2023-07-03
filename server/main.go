package main

import (
	"context"
	"fmt"
	"time"
)

type Server interface {
	Start()
	Stop()
}

func main() {
	var s Serv
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	RunServer(ctx, s)
	cancel()
}

type Serv struct{}

func (serv Serv) Start() {
	fmt.Println("Start")
}

func (serv Serv) Stop() {
	fmt.Println("Stop")
}

func RunServer(ctx context.Context, serv Server) {
	ch := make(chan struct{})
	serv.Start()
	go func(ctx context.Context) {
		<-ctx.Done()
		serv.Stop()
		ch <- struct{}{}
	}(ctx)
	<-ch
}
