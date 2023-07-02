package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go Infinite(1*time.Second, wg)

	wg.Add(1)
	go Infinite(3*time.Second, wg)

	wg.Wait()

}

func Infinite(t time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("adsadadasdasd")
	time.Sleep(t)
}
