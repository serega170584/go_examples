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

type ServerInst struct{}

func (server ServerInst) Start() {
	fmt.Println("Start")
}

func (server ServerInst) Stop() {
	fmt.Println("Stop")
}

func Run(ctx context.Context, server Server) {
	ch := make(chan struct{})
	server.Start()
	go func() {
		<-ctx.Done()
		server.Stop()
		ch <- struct{}{}
	}()
	<-ch
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	Run(ctx, ServerInst{})
}
