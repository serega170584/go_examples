package main

import "fmt"

func main() {
	ch := make(chan struct{})
	cnt := 10
	res := make(chan int, cnt)

	for i := 0; i < cnt; i++ {
		i := i
		go func() {
			res <- i
			<-ch
		}()
	}

	go func() {
		for {
			ch <- struct{}{}
			cnt--
			if cnt == 0 {
				close(res)
			}
		}
	}()

	for val := range res {
		fmt.Println(val)
	}
}
