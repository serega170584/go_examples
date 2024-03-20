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

	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	fmt.Println(getMaxWin(n, list, a, b, k))
}

func getMaxWin(n int, list []int, a int, b int, k int) int {
	maxNumber := 0

	minSearchCnt := 0
	if a-1 >= k {
		minSearchCnt = (a - 1) / k
	}

	maxSearchCnt := 0
	if b-1 >= k {
		maxSearchCnt = (b - 1) / k
	}

	if maxSearchCnt-minSearchCnt >= n-1 {
		for i := 0; i < n; i++ {
			maxNumber = max(maxNumber, list[i])
		}
		return maxNumber
	}

	minSearchInd := minSearchCnt % n
	maxSearchInd := maxSearchCnt % n

	if minSearchInd <= maxSearchInd {
		for i := minSearchInd; i <= maxSearchInd; i++ {
			maxNumber = max(maxNumber, list[i])
		}
	} else {
		for i := minSearchInd; i < n; i++ {
			maxNumber = max(maxNumber, list[i])
		}

		for i := 0; i <= maxSearchInd; i++ {
			maxNumber = max(maxNumber, list[i])
		}
	}

	minSearchInd = 0
	if maxSearchCnt != 0 {
		minSearchInd = n - maxSearchCnt%n
	}

	maxSearchInd = 0
	if minSearchCnt != 0 {
		maxSearchInd = n - minSearchCnt%n
	}

	if minSearchInd <= maxSearchInd {
		for i := minSearchInd; i <= maxSearchInd; i++ {
			maxNumber = max(maxNumber, list[i])
		}
	} else {
		for i := minSearchInd; i < n; i++ {
			maxNumber = max(maxNumber, list[i])
		}

		for i := 0; i < maxSearchInd; i++ {
			maxNumber = max(maxNumber, list[i])
		}
	}

	return maxNumber
}
