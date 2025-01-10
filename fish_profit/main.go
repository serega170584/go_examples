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

	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	maxProfit := 0
	for i := range list {
		if i+k == n {
			break
		}
		for j := i + 1; j <= i+k; j++ {
			maxProfit = max(maxProfit, list[j]-list[i])
		}
	}

	fmt.Println(maxProfit)
}
