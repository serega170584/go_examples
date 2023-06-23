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
		v := q[curIndex]
		l := v.left
		r := v.right

		curIndex++

		if l >= r {
			continue loop
		}

		if l == r-1 && list[l] > list[r] {
			list[l], list[r] = list[r], list[l]
			continue loop
		}

		if l == r-1 {
			continue loop
		}

		fix := l
		l++

		for l <= r {
			for l <= v.right && list[l] <= list[fix] {
				l++
			}

			for r > fix && list[fix] <= list[r] {
				r--
			}

			if l < r {
				list[l], list[r] = list[r], list[l]
				l++
				r--
			}
		}

		list[fix], list[r] = list[r], list[fix]

		q = append(q, lr{left: fix, right: r - 1})
		counter++

		q = append(q, lr{left: r + 1, right: v.right})
		counter++
	}

	return list
}

type lr struct {
	left  int
	right int
}
