package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter rows count")
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter columns count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list")
	list := make([][]int, m)
	for i := 0; i < m; i++ {
		list[i] = make([]int, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			list[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}

	fmt.Println("Got turtle min x without debt", minWithoutDebtX(m, n, list))
}

func turtleMaxCost(m int, n int, list [][]int) int {
	maxList := make([][]int, m)
	for i := 0; i < m; i++ {
		maxList[i] = make([]int, n)
		for j := 0; j < n; j++ {
			maxList[i][j] = -math.MaxInt
		}
	}

	maxList[0][0] = list[0][0]

	for i := range list {
		for j := range list[i] {
			prevI := i - 1
			prevJ := j - 1
			if prevI >= 0 && maxList[prevI][j] >= 0 {
				maxList[i][j] = max(maxList[prevI][j]+list[i][j], maxList[i][j])
			}
			if prevJ >= 0 && maxList[i][prevJ] >= 0 {
				maxList[i][j] = max(maxList[i][prevJ]+list[i][j], maxList[i][j])
			}
		}
	}
	return maxList[m-1][n-1]
}

func minWithoutDebtX(m int, n int, list [][]int) int {
	minPath := 0
	maxPath := 0
	for _, row := range list {
		for _, v := range row {
			if v > 0 {
				maxPath += v
			}
			if v < 0 {
				minPath += v
			}
		}
	}
	l := -maxPath
	r := -minPath

	for l < r {
		mid := (l + r) / 2
		if check(mid, m, n, list) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func check(mid int, m int, n int, list [][]int) bool {
	list[0][0] = mid
	cost := turtleMaxCost(m, n, list)
	if cost < 0 {
		return true
	}
	return false
}
