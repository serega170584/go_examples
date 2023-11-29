package main

import "fmt"

func main() {
	fmt.Println(search([]int{1, 3, 6, 9}, 5))
}

func search(slice []int, val int) int {
	left, right := 0, len(slice)-1
	var middle int
	for left <= right {
		middle = (left + right) / 2
		if slice[middle] < val {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}
