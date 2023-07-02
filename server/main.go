package main

import "context"

type Server interface {
	Start()
	Stop()
}

func main() {
	var s Server
	ctx := context.Background()
	go func(ctx context.Context) {
		<-ctx.Done()
		s.Stop()
	}(ctx)
}
