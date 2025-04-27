package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	a := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	b := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}

	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())

	dp := make([][]int, cnt+1)
	for i := 0; i < cnt+1; i++ {
		dp[i] = make([]int, c+1)
	}

	for i := 1; i < cnt+1; i++ {
		for j := 1; j < c+1; j++ {
			dp[i][j] = dp[i-1][j]
			if b[i-1] <= j {
				if dp[i][j] < dp[i-1][j-b[i-1]]+a[i-1] {
					dp[i][j] = dp[i-1][j-b[i-1]] + a[i-1]
				}
			}
		}
	}

	fmt.Println(dp[cnt][c])
}
