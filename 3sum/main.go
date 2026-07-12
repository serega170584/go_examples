package main

import (
	"fmt"
	"sort"
)

func main() {
	// -4 -1 -1 -1 0 1 2
	nums := []int{-1, -1, -1, -1, -1, 0, 1, 1, 1, 2, -4}
	sort.Ints(nums)

	res := make([][3]int, 0)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}

			if right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}

			if nums[i]+nums[left]+nums[right] < 0 {
				left++
				continue
			}

			if nums[i]+nums[left]+nums[right] > 0 {
				right--
				continue
			}

			res = append(res, [3]int{nums[i], nums[left], nums[right]})
			left++
			right--
		}
	}

	fmt.Println(res)
}
