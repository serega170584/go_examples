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

	dp := make([][][]bool, n)

	for i := 0; i < n; i++ {
		dp[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]bool, n+1)
			if a[i] == b[j] {
				dp[i][j][1] = true
			}
		}
	}

	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			for j := 0; j <= n-l; j++ {
				for length := 1; length < l; length++ {
					if dp[i][j][length] && dp[i+length][j+length][l-length] {
						dp[i][j][l] = true
						break
					}

					if dp[i][j+l-length][length] && dp[i+length][j][l-length] {
						dp[i][j][l] = true
						break
					}
				}
			}
		}
	}

	fmt.Println(dp[0][0][n])
}
