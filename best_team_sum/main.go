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

	fmt.Println("Enter count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list")
	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Got best team sum", bestTeamSum(n, list))
}

func bestTeamSum(n int, list []int) int {
	bestSum := 0
	curSum := 0
	r := 0
	for l := range list {
		for r < n && (l == r || list[l]+list[l+1] >= list[r]) {
			curSum += list[r]
			r++
		}
		bestSum = max(bestSum, curSum)
		curSum -= list[l]
	}
	return bestSum
}
