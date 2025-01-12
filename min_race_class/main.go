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

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	cost := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			scanner.Scan()
			v, _ := strconv.Atoi(scanner.Text())
			cost[i] = append(cost[i], v)
		}
	}

	rowDirectMax := make([]int, n)
	rowInDirectMax := make([]int, n)
	for i := 0; i < n; i++ {
		rowDirectMax[i] = cost[i][0]
		rowInDirectMax[i] = cost[i][1]
		if rowInDirectMax[i] > rowDirectMax[i] {
			rowDirectMax[i], rowInDirectMax[i] = rowInDirectMax[i], rowDirectMax[i]
		}
		for j := 2; j < m; j++ {
			if cost[i][j] > rowDirectMax[i] {
				rowDirectMax[i], rowInDirectMax[i] = cost[i][j], rowDirectMax[i]
				continue
			}
			if cost[i][j] > rowInDirectMax[i] {
				rowInDirectMax[i] = cost[i][j]
			}
		}
	}

	colRestMax := make([][]int, m)
	for i := 0; i < m; i++ {
		colRestMax[i] = make([]int, n)
		for j := 0; j < n; j++ {
			colRestMax[i][j] = rowDirectMax[j]
			if cost[j][i] == colRestMax[i][j] {
				colRestMax[i][j] = rowInDirectMax[j]
			}
		}
	}

	downRowForColRestMax := make([][]int, n)
	downRowForColRestMax[0] = make([]int, m)
	downRowForColRestMax[1] = make([]int, m)
	for j := 0; j < m; j++ {
		downRowForColRestMax[1][j] = colRestMax[j][0]
	}
	for i := 2; i < n; i++ {
		downRowForColRestMax[i] = make([]int, m)
		for j := 0; j < m; j++ {
			downRowForColRestMax[i][j] = downRowForColRestMax[i-1][j]
			if colRestMax[j][i-1] > downRowForColRestMax[i][j] {
				downRowForColRestMax[i][j] = colRestMax[j][i-1]
			}
		}
	}

	upRowForColRestMax := make([][]int, n)
	for i := 0; i < n; i++ {
		upRowForColRestMax[i] = make([]int, m)
	}
	for j := 0; j < m; j++ {
		upRowForColRestMax[n-2][j] = colRestMax[j][n-1]
	}
	for i := n - 3; i >= 0; i-- {
		for j := 0; j < m; j++ {
			upRowForColRestMax[i][j] = upRowForColRestMax[i+1][j]
			if colRestMax[j][i+1] > upRowForColRestMax[i][j] {
				upRowForColRestMax[i][j] = colRestMax[j][i+1]
			}
		}
	}

	resRow := 0
	resCol := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if downRowForColRestMax[i][j] < upRowForColRestMax[i][j] {
				downRowForColRestMax[i][j] = upRowForColRestMax[i][j]
				resRow = i
				resCol = j
			}
		}
	}

	minMax := math.MaxInt
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if downRowForColRestMax[i][j] < minMax {
				minMax = downRowForColRestMax[i][j]
				resRow = i
				resCol = j
			}
		}
	}

	fmt.Println(resRow+1, resCol+1)
}
