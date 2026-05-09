package main

import (
	"fmt"
	"sort"
)

func main() {
	// -4 -1 -1 -1 0 1 2
	nums := []int{-1, -1, -1, -1, -1, 0, 1, 1, 1, 2, -4}
	sort.Ints(nums)

	res := [][3]int{}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		l := i + 1
		r := len(nums) - 1
		for l < r {
			if l-1 != i && nums[l-1] == nums[l] {
				l++
				continue
			}

			if r != len(nums)-1 && nums[r+1] == nums[r] {
				r--
				continue
			}

			s := nums[i] + nums[l] + nums[r]

			if s < 0 {
				l++
				continue
			}

			if s > 0 {
				r--
				continue
			}

			if s == 0 {
				res = append(res, [3]int{nums[i], nums[l], nums[r]})
				l++
				r--
			}
		}
	}

	fmt.Println(res)

}
