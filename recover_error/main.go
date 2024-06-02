package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make([]int, 0)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer handler()
		defer wg.Done()
		fmt.Println(a)
		//a[1] = 123
	}()

	wg.Wait()
}

func handler() {
	e := recover()
	//err, _ := e.(error)
	fmt.Println(e)
}
