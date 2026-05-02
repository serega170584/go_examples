package main

import "fmt"

func main() {
	str := "abba"
	v := []rune(str)
	n := 4
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	maxPL := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			left := j
			right := i + j

			if left == right {
				dp[left][right] = 1
				continue
			}

			if right-left == 1 && v[left] == v[right] {
				dp[left][right] = 2
			}

			if dp[left][right] > maxPL {
				maxPL = dp[left][right]
				continue
			}

			if dp[left+1][right-1] == 0 {
				continue
			}

			if v[left] != v[right] {
				dp[left][right] = 0
				continue
			}

			dp[left][right] = dp[left+1][right-1] + 2

			if dp[left][right] > maxPL {
				maxPL = dp[left][right]
			}
		}
	}

	fmt.Println(maxPL)
}
