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
	m, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	list := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			scanner.Scan()
			v, _ := strconv.Atoi(scanner.Text())
			list[i] = append(list[i], v)
		}
	}

	fmt.Println(getMaxDenyPoint(list, m, n))
}

func getMaxDenyPoint(list [][]int, m int, n int) (int, int) {
	rowMax1 := make([]int, m)
	rowMax2 := make([]int, m)

	for i, row := range list {
		for _, v := range row {
			if v > rowMax1[i] {
				rowMax2[i] = rowMax1[i]
				rowMax1[i] = v
			} else if v > rowMax2[i] {
				rowMax2[i] = v
			}
		}
	}

	withExcludedMax := make([][]int, m)
	for i, row := range list {
		withExcludedMax[i] = make([]int, n)
		for j, v := range row {
			withExcludedMax[i][j] = rowMax1[i]
			if v == rowMax1[i] {
				withExcludedMax[i][j] = rowMax2[i]
			}
		}
	}

	excludedMax := make([][]int, m)
	excludedMax[0] = make([]int, n)

	for i := 1; i < m; i++ {
		excludedMax[i] = make([]int, n)
		for j := 0; j < n; j++ {
			excludedMax[i][j] = max(excludedMax[i-1][j], withExcludedMax[i-1][j])
		}
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			excludedMax[i][j] = max(excludedMax[i+1][j], withExcludedMax[i+1][j])
		}
	}

	minMaxVal := excludedMax[0][0]
	rowIndex := 1
	colIndex := 1
	for i, row := range excludedMax {
		for j, v := range row {
			if v < minMaxVal {
				minMaxVal = v
				rowIndex = i + 1
				colIndex = j + 1
			}
		}
	}

	return rowIndex, colIndex

}
