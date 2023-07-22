package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {
	var cnt int
	fmt.Println("Enter count")
	_, err := fmt.Scanln(&cnt)

	if cnt < 2 {
		log.Fatal(errors.New("Invalid count"))
	}

	if err != nil {
		log.Fatal(err)
	}

	x := make([]interface{}, cnt)
	nails := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		x[i] = &nails[i]
	}

	fmt.Println("Enter nails")
	_, err = fmt.Scanln(x...)
	if err != nil {
		log.Fatal(err)
	}

	nails = sort(nails)

	fmt.Printf("Min length: %d", minLength(nails))

}

func minLength(nails []int) int {
	var minLength int
	cnt := len(nails)

	weights := make([]int, cnt-1)

	if len(nails) == 2 {
		return nails[1] - nails[0]
	}

	for i := 1; i < cnt; i++ {
		weights[i-1] = nails[i] - nails[i-1]
	}

	weights = append(weights, math.MaxInt)

	cmpIdx := 0
	minLength, prevLength := weights[0], weights[0]
loop:
	for i := 1; i < cnt; i++ {
		curLength := weights[i]

		if i == cmpIdx+1 {
			prevLength = curLength
			continue loop
		}

		if curLength < prevLength {
			cmpIdx = i
			minLength += curLength
		} else {
			minLength += prevLength
		}

		prevLength = curLength
	}

	return minLength
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
