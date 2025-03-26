package main

import (
	"bufio"
	"fmt"
	"os"
)

//    ""  a b c
// ""  1  1 1 1
// d   0  0 0 0
// c   0  0 0 0

//    ""  a b c
// ""  1  1 1 1
// c   0  0 0 1
// d   0  0 0 0

// aba
// ab_
// abab
// ab__

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	s := []rune(scanner.Text())

	l := len(s) + 1
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}

	for i := 1; i < l; i++ {
		for j := 1; j < l; j++ {
			if s[i-1] == s[j-1] && i != j {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i-1][j]
				if dp[i][j-1] > dp[i][j] {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	if dp[l-1][l-1] >= 2 {
		fmt.Println(1)
		return
	}

	fmt.Println(0)
}
