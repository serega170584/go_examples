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

	var a [2][]int
	a[0] = make([]int, cnt)
	a[1] = make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		a[0][i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		a[1][i], _ = strconv.Atoi(scanner.Text())
	}

	var dp [2][]int
	dp[0] = make([]int, cnt)
	dp[1] = make([]int, cnt)

	dp[0][0] = a[0][0]
	dp[1][0] = a[1][0]

	maxVal := dp[0][0]
	if maxVal < dp[1][0] {
		maxVal = dp[1][0]
	}

	if cnt == 1 {
		fmt.Println(maxVal)
		return
	}

	dp[0][1] = a[0][1]
	dp[1][1] = a[1][1]

	if maxVal < dp[0][1] {
		maxVal = dp[0][1]
	}

	if maxVal < dp[1][1] {
		maxVal = dp[1][1]
	}

	if cnt == 2 {
		fmt.Println(maxVal)
	}

	for i := 2; i < cnt; i++ {
		for j := 0; j < i-1; j++ {
			if dp[0][i] < dp[0][j]+a[0][i] {
				dp[0][i] = dp[0][j] + a[0][i]
			}

			if dp[0][i] < dp[1][j]+a[0][i] {
				dp[0][i] = dp[1][j] + a[0][i]
			}

			if dp[1][i] < dp[0][j]+a[1][i] {
				dp[1][i] = dp[0][j] + a[1][i]
			}

			if dp[1][i] < dp[1][j]+a[1][i] {
				dp[1][i] = dp[1][j] + a[1][i]
			}
		}
		if maxVal < dp[0][i] {
			maxVal = dp[0][i]
		}
		if maxVal < dp[1][i] {
			maxVal = dp[1][i]
		}
	}

	fmt.Println(maxVal)
}
