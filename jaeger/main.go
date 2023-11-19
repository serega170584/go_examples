package main

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"log"
	"time"
)

func main() {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeRateLimiting,
			Param: 100,
		},
		ServiceName: "test application",
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Printf("Error")
	}
	defer func(closer io.Closer) {
		err := closer.Close()
		if err != nil {
			log.Printf("Closer error")
		}
	}(closer)

	opentracing.SetGlobalTracer(tracer)

	span, ctx := opentracing.StartSpanFromContext(context.Background(), "Operation 123")
	defer span.Finish()

	//go func(tracer opentracing.Tracer) {
	//	opentracing.SetGlobalTracer(tracer)
	//	span, _ := opentracing.StartSpanFromContext(context.Background(), "Goroutine operation")
	//	defer span.Finish()
	//	for i := 0; i < 6; i++ {
	//		time.Sleep(400 * time.Millisecond)
	//		println("test")
	//	}
	//}(tracer)

	func(ctx context.Context) {
		span, _ := opentracing.StartSpanFromContext(ctx, "Operation 456")
		defer span.Finish()
		for i := 0; i < 3; i++ {
			time.Sleep(400 * time.Millisecond)
		}
	}(ctx)

	for i := 0; i < 3; i++ {
		time.Sleep(3 * time.Second)
	}
}
