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

	al := len(a)
	bl := len(b)
	dp := make([][]int, al+1)

	dp[0] = make([]int, bl+1)

	for j := 0; j <= bl; j++ {
		dp[0][j] = j
	}

	for i := 0; i <= al; i++ {
		dp[i] = make([]int, bl+1)
		dp[i][0] = i
	}

	for i := 1; i <= al; i++ {
		for j := 1; j <= bl; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}

			dp[i][j] = dp[i-1][j-1] + 1

			if dp[i-1][j]+1 < dp[i][j] {
				dp[i][j] = dp[i-1][j] + 1
			}

			if dp[i][j-1]+1 < dp[i][j] {
				dp[i][j] = dp[i][j-1] + 1
			}
		}
	}

	fmt.Println(dp[al][bl])
}
