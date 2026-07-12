package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 8, 8, 10}
	target := 8

	l := 0
	r := len(nums) - 1
	for l < r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	first := l

	l = 0
	r = len(nums) - 1

	for l < r {
		m := (l + r + 1) / 2
		if nums[m] <= target {
			l = m
		} else {
			r = m - 1
		}
	}
	second := l

	fmt.Println(first, second)
}
