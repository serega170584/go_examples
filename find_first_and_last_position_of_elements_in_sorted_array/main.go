package main

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8

	l := 0
	r := len(nums) - 1
	for l < r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
			continue
		}

		r = m
	}

	ls := l

}
