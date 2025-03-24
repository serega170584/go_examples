package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	s := []rune(scanner.Text())

	n := len(s)
	dp := make([][]rune, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]rune, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				dp[i][j] = 1
				continue
			}

			if s[i] == s[j] {
				if i == j+1 {
					dp[i][j] = 2
				} else {
					dp[i][j] = dp[i+1][j-1] + 2
				}
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	fmt.Println(dp[0][n-1])
}
