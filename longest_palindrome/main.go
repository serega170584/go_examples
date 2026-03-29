package main

import (
	"fmt"
)

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	sl := []rune(s)
	dp := make([][]int, len(sl))
	for i := 0; i < len(sl); i++ {
		dp[i] = make([]int, len(sl))
		dp[i][i] = 1
	}

	res := 1
	maxLeft := 0
	maxRight := 0

	for k := 1; k < len(sl); k++ {
		for i := 0; i < len(sl)-k; i++ {
			j := i + k
			if j-i == 1 && sl[i] == sl[j] {
				dp[i][j] = 2
			} else if sl[i] == sl[j] && dp[i+1][j-1] != 0 {
				dp[i][j] = dp[i+1][j-1] + 2
			}

			if res < dp[i][j] {
				res = dp[i][j]
				maxLeft = i
				maxRight = j
			}
		}
	}

	sRes := ""
	for i := maxLeft; i <= maxRight; i++ {
		sRes += string(sl[i])
	}

	return sRes
}
