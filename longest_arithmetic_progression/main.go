package main

import "fmt"

type ex interface {
	ex()
}

type al int

func (a *al) ex() {
	fmt.Println("123")
}

func ex1(a ex) {
	a.ex()
}

func main() {
	var a al
	ex1(&a)
}

func lengthOfLongestAP(A []int) int {
	if len(A) < 2 {
		return len(A)
	}

	dp := make([]map[int]int, len(A))
	for i := 0; i < len(dp); i++ {
		dp[i] = make(map[int]int)
	}

	longest := 2
	for i := 1; i < len(A); i++ {
		for j := 0; j < i; j++ {
			if _, ok := dp[j][A[i]-A[j]]; ok {
				dp[i][A[i]-A[j]] = dp[j][A[i]-A[j]] + 1
				if dp[i][A[i]-A[j]] > longest {
					longest = dp[i][A[i]-A[j]]
				}
				continue
			}
			dp[i][A[i]-A[j]] = 2
		}
	}

	return longest
}
