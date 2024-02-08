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

	fmt.Println("Enter list count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list")
	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Enter number")
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Got pairs count ", cntPairsWithDiffGtk(n, list, k))
}

func cntPairsWithDiffGtk(cnt int, list []int, k int) int {
	pairsCnt := 0
	left := 0
	right := 0
	for range list {
		for right < cnt && list[right]-list[left] <= k {
			right++
		}
		pairsCnt += cnt - right
		left++
	}
	return pairsCnt
}
