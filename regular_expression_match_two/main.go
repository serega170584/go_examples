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
	a := []rune(scanner.Text())

	scanner.Scan()
	b := []rune(scanner.Text())

	n := len(a)
	m := len(b)

	dp := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true
	for j := 1; j < m+1; j++ {
		if b[j-1] == '*' && j > 1 {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if a[i-1] == b[j-1] || b[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if b[j-1] == '*' && j > 1 {
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (b[j-2] == a[i-1] || b[j-2] == '.'))
			}
		}
	}

	fmt.Println(dp[n][m])
}
