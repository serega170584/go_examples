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

	al := len(a) + 1
	bl := len(b) + 1

	dp := make([][]int, bl)
	dp[0] = make([]int, al)
	for j := 0; j < al; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < bl; i++ {
		dp[i] = make([]int, al)
	}

	for i := 1; i < bl; i++ {
		for j := 1; j < al; j++ {
			if b[i-1] == a[j-1] {
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	fmt.Println(dp[bl-1][al-1])
}
