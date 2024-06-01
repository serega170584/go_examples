package main

import "fmt"

func main() {
	fmt.Println(getLongestSequenceSize([]int{1, 5, 2, 6, 3}))
}

func getLongestSequenceSize(list []int) int {
	if len(list) == 0 {
		return 0
	}

	increasedSequenceMaxSize := 1
	decreasedSequenceMaxSize := 1
	increasedSequenceSize := 1
	decreasedSequenceSize := 1
	prev := list[0]

	for i := 1; i < len(list); i++ {
		if list[i] > prev {
			increasedSequenceSize++
		} else {
			increasedSequenceMaxSize = max(increasedSequenceMaxSize, increasedSequenceSize)
			increasedSequenceSize = 1
		}

		if list[i] < prev {
			decreasedSequenceSize++
		} else {
			decreasedSequenceMaxSize = max(decreasedSequenceMaxSize, decreasedSequenceSize)
			decreasedSequenceSize = 1
		}

		prev = list[i]
	}

	increasedSequenceMaxSize = max(increasedSequenceMaxSize, increasedSequenceSize)
	decreasedSequenceMaxSize = max(decreasedSequenceMaxSize, decreasedSequenceSize)

	return max(increasedSequenceMaxSize, decreasedSequenceMaxSize)
}
