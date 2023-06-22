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

	list = quickSort(list)
	fmt.Println(list)
}

func quickSort(list []int) []int {
	var curIndex int
	counter := 1
	cnt := len(list)

	q := []lr{{left: 0, right: cnt - 1}}
loop:
	for counter > 0 {
		counter--
		curLR := q[curIndex]
		left := curLR.left
		right := curLR.right

		curIndex++

		if left >= right {
			continue loop
		}

		if left == right-1 && list[left] > list[right] {
			list[left], list[right] = list[right], list[left]
			continue loop
		}

		if left == right-1 {
			continue loop
		}

		fix := left
		left++

		for left <= right {
			for left <= curLR.right && list[left] <= list[fix] {
				left++
			}

			for right > fix && list[fix] <= list[right] {
				right--
			}

			if left < right {
				list[left], list[right] = list[right], list[left]
				left++
				right--
			}
		}

		list[fix], list[right] = list[right], list[fix]

		q = append(q, lr{left: fix, right: right - 1})
		counter++

		q = append(q, lr{left: right + 1, right: curLR.right})
		counter++
	}

	return list
}

type lr struct {
	left, right int
}
