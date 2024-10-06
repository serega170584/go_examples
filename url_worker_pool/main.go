package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	defer func() {
		fmt.Println(time.Since(t))
	}()

	in := make(chan string, 5)
	url := [5]string{
		"https://ya.ru",
		"https://magnit.ru",
		"https://ozon.ru",
		"https://avito.ru",
		"https://tinkoff.ru",
	}
	for i := 0; i < 5; i++ {
		in <- url[i]
	}
	close(in)

	out := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				resp, err := http.Get(v)
				if err != nil {
					out <- err.Error()
					continue
				}

				if resp.StatusCode != http.StatusOK {
					out <- "Error"
					continue
				}

				out <- "Success"
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for v := range out {
		fmt.Println(v)
	}
}
