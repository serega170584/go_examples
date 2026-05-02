package main

import (
	"fmt"
	"time"
)

func main() {

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("123")
	}

	//var mu sync.RWMutex
	//go func() {
	//	time.Sleep(3 * time.Second)
	//	mu.RUnlock()
	//	fmt.Println("End 123")
	//}()
	//mu.RLock()
	//fmt.Println("123")
	//mu.RLock()
	//fmt.Println("456")
	//mu.RUnlock()
	//
	//t := time.Now()
	//defer func() {
	//	fmt.Println(time.Since(t))
	//}()
	//
	//var cnt uint64
	//var wg sync.WaitGroup
	//wg.Add(10000)
	//for i := 0; i < 10000; i++ {
	//	go func() {
	//		defer wg.Done()
	//		atomic.AddUint64(&cnt, 1)
	//	}()
	//}
	//
	//wg.Wait()
	//
	//fmt.Println(cnt)
	//
	//errgroup.WithContext(ctx)
	//
	//var g errgroup.Group
	//
	//var urls = []string{
	//	"http://www.google1111.com",
	//	"http://www.google.com", // This will likely return an error
	//	"http://www.golang.org",
	//}
	//
	//g.Go(func() error {
	//	for _, v := range urls {
	//		resp, err := http.Get(v)
	//		if err != nil {
	//			return err
	//		}
	//		fmt.Println("")
	//	}
	//})
	//
	//fmt.Println("Successfully fetched all URLs.")
}
