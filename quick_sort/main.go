package main

import (
	"fmt"
	"math/rand"
)

func main() {
	cnt := 20
	list := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		list[i] = rand.Intn(9)
	}

	list = sort(list)
	fmt.Println(list)
}

type lr struct {
	left  int
	right int
}

func sort(list []int) []int {
	cnt := len(list)

	q := []lr{{left: 0, right: cnt - 1}}
	curIndex := 0
	counter := 1

	//list = []int{669, 65, 838, 579, 324, 266, 801, 912, 823, 249, 670, 659, 452, 421, 886, 598, 956, 905, 929, 77}

	fmt.Println(list)

loop:
	for counter != 0 {
		counter--

		val := q[curIndex]
		fmt.Println(val)
		fmt.Println(list)

		if val.left >= val.right {
			curIndex++
			continue loop
		}

		if val.left == val.right-1 && list[val.left] > list[val.right] {
			list[val.left], list[val.right] = list[val.right], list[val.left]
			curIndex++
			continue loop
		}

		if val.left == val.right-1 {
			curIndex++
			continue loop
		}

		fix := val.left
		left := val.left + 1
		right := val.right

		for left <= right {
			for left <= val.right && list[fix] >= list[left] {
				left++
			}

			for right > fix && list[right] >= list[fix] {
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

		q = append(q, lr{left: right + 1, right: val.right})
		counter++

		curIndex++
	}

	fmt.Println(q)

	return list
}
