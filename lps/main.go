package main

import "fmt"

//	a b c
//
// a 1 0
// b   1 0
// c     1
func main() {
	a := []rune("cbbd")
	mpl := 0
	dp := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		dp[i] = make([]int, len(a))
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i; j++ {
			l := 0
			left := j
			right := j + i
			if left == right {
				l = 1
			}

			if right-left == 1 && a[left] == a[right] {
				l = 2
			}

			if right-left > 1 && dp[left+1][right-1] != 0 && a[left] == a[right] {
				l = dp[left+1][right-1] + 2
			}
			dp[left][right] = l

			if l > mpl {
				mpl = l
			}
		}
	}

	fmt.Println(mpl)
}
