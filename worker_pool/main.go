package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const workersCnt = 3

func main() {
	urls := []string{
		"http://yandex.ru",
		"http://magnit.ru",
		"http://ozon.ru",
		"http://avito.ru",
	}
	in := make(chan string, len(urls))
	out := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(workersCnt)

	go func(in chan string, urls []string) {
		defer close(in)
		for _, v := range urls {
			in <- v
		}
	}(in, urls)

	go func(ctx context.Context, in chan string, out chan string, wg *sync.WaitGroup) {
		for i := 0; i < workersCnt; i++ {
			go worker(ctx, in, out, wg)
		}
	}(ctx, in, out, &wg)

	go func(wg *sync.WaitGroup, out chan string) {
		defer close(out)
		wg.Wait()
	}(&wg, out)

	for v := range out {
		fmt.Println(v)
	}
}

func worker(ctx context.Context, in chan string, out chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range in {
		sig := make(chan string)
		adapter(ctx, out, v, sig)
	}
}

func handleUrl(url string) string {
	rsp, err := http.Get(url)
	if err != nil || rsp.StatusCode != 200 {
		return fmt.Sprintf("Url: %s, error", url)
	}
	return fmt.Sprintf("Url: %s, OK", url)
}

func adapter(ctx context.Context, out chan string, url string, sig chan string) {
	go func(sig chan string, url string) {
		sig <- handleUrl(url)
	}(sig, url)

	select {
	case <-ctx.Done():
		out <- ctx.Err().Error()
	case res := <-sig:
		out <- res
	}
}
