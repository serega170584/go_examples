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

func (s ServerInst) Start() {
	fmt.Println("Start")
}

func (s ServerInst) Stop() {
	fmt.Println("Stop")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	ch := make(chan struct{})
	s := ServerInst{}
	Run(ctx, s, ch)
	<-ch
}

func Run(ctx context.Context, s Server, ch chan struct{}) {
	s.Start()
	go func() {
		<-ctx.Done()
		s.Stop()
		ch <- struct{}{}
	}()
}
