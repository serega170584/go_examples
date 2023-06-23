package main

import (
	"fmt"
	"math/rand"
)

func main() {
	cnt := 20
	list := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		list[i] = rand.Intn(1000)
	}

	fmt.Println(pyramidSort(list))
}

func pyramidSort(list []int) []int {

	li := len(list) - 1
	for li > 0 {
	loop:
		for i := range list {
			if i > li {
				break loop
			}

			pi := i
			for pi != 0 {
				ppi := pi/2 + pi%2 - 1
				if list[ppi] < list[pi] {
					list[pi], list[ppi] = list[ppi], list[pi]
				}
				pi = ppi
			}
		}
		list[0], list[li] = list[li], list[0]
		li--
	}

	return list
}
