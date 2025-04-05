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

	a := make([]int, 0, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		a = append(a, v)
	}

	if cnt == 1 {
		fmt.Println(0)
	}

	dp := make([][][2]int, cnt)
	for i := 0; i < cnt; i++ {
		dp[i] = make([][2]int, cnt)
		dp[i][i] = [2]int{a[i], 0}
	}

	for offset := 1; offset < cnt; offset++ {
		for i := 0; i < cnt-offset; i++ {
			last := i + offset
			for j := i; j < last; j++ {
				dp[i][last][0] = dp[i][j][0] + dp[j+1][last][0]
				if dp[i][last][1] == 0 || dp[i][last][1] > dp[i][j][1]+dp[j+1][last][1]+dp[i][last][0] {
					dp[i][last][1] = dp[i][j][1] + dp[j+1][last][1] + dp[i][last][0]
				}
			}
		}
	}

	fmt.Println(dp[0][cnt-1][1])

}
