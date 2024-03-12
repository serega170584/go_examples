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
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	prices := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		prices[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(fishProfit(prices, n, k))
}

func fishProfit(prices []int, n int, k int) int {
	maxProfit := 0
	for i := 0; i < n-1; i++ {
		last := i + k + 1
		if i+k+1 > n {
			last = n
		}
		for j := i + 1; j < last; j++ {
			maxProfit = max(maxProfit, prices[j]-prices[i])
		}
	}
	return maxProfit
}
