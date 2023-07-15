package main

import "fmt"

func main() {
	a := []int{1, 3, 4, 7, 10, 14, 17, 20, 23, 30}
	b := []int{2, 7, 8, 9, 13, 15, 16, 21, 25, 27, 28, 29, 31, 33, 34}
	cnt := len(a) + len(b)
	chanA := make(chan int)
	chanB := make(chan int)
	chanMerge := make(chan int)
	go func() {
		defer close(chanA)
		for _, val := range a {
			chanA <- val
		}
	}()

	go func() {
		defer close(chanB)
		for _, val := range b {
			chanB <- val
		}
	}()

	go func() {
		valA, okA := <-chanA
		valB, okB := <-chanB
	loop:
		for i := 0; i < cnt; i++ {
			if !(okA || okB) {
				return
			}

			if !okA {
				chanMerge <- valB
				valB, okB = <-chanB
				continue loop
			}

			if !okB {
				chanMerge <- valA
				valA, okA = <-chanA
				continue loop
			}

			if valA <= valB {
				chanMerge <- valA
				valA, okA = <-chanA
				continue loop
			}

			if valA > valB {
				chanMerge <- valB
				valB, okB = <-chanB
				continue loop
			}
		}
		close(chanMerge)
	}()

	for val := range chanMerge {
		fmt.Println(val)
	}
}
