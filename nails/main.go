package main

import (
	"fmt"
	"log"
)

// 1 2 3
// 1 2 3 4
// 1 2 5 4
// 1 2
// 1 3 5 6
// 1 3 5 10
// 1 3 5 7 12
// 1 3 5 7 9 14
// 1 4 7 10 12 15
// 1 4 7 10 12 14
// 1 4 7 10 13 15 17
func main() {
	var cnt int
	_, err := fmt.Scanln(&cnt)

	if err != nil {
		log.Fatal(err)
	}

	x := make([]interface{}, cnt)
	nails := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		x[i] = &nails[i]
	}

	_, err = fmt.Scanln(x...)
	if err != nil {
		log.Fatal(err)
	}

	nails = sort(nails)
	fmt.Println(nails)

	fmt.Printf("%d", minLength(nails))
}

func minLength(nails []int) int {
	cnt := len(nails)

	if len(nails) == 2 {
		return nails[1] - nails[0]
	}

	if len(nails) == 3 {
		return nails[2] - nails[0]
	}

	if len(nails) == 4 {
		return nails[3] - nails[2] + nails[1] - nails[0]
	}

	minLength := make([]int, cnt)
	minLength[2] = nails[2] - nails[0]
	minLength[3] = nails[3] - nails[2] + nails[1] - nails[0]
	for i := 4; i < cnt; i++ {
		if minLength[i-1] < minLength[i-2] {
			minLength[i] = minLength[i-1] + nails[i] - nails[i-1]
		} else {
			minLength[i] = minLength[i-2] + nails[i] - nails[i-1]
		}
	}
	return minLength[cnt-1]
}

func sort(nails []int) []int {

	lastIdx := len(nails) - 1
	for lastIdx != 0 {
		for idx := 0; idx <= lastIdx; idx++ {

			pyramidIdx := idx

			for pyramidIdx != 0 {
				parentIdx := pyramidIdx/2 + pyramidIdx%2 - 1
				if nails[pyramidIdx] > nails[parentIdx] {
					nails[pyramidIdx], nails[parentIdx] = nails[parentIdx], nails[pyramidIdx]
				}
				pyramidIdx = parentIdx
			}
		}

		nails[0], nails[lastIdx] = nails[lastIdx], nails[0]

		lastIdx--
	}

	return nails
}
