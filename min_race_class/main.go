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

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			scanner.Scan()
			a[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}

	fmt.Println(maxDeny(n, m, a))
}

func maxDeny(n int, m int, a [][]int) (int, int) {
	minExcludeMaxVal := math.MaxInt
	row := -1
	col := -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			val := getExcludeMaxVal(i, j, n, m, a, minExcludeMaxVal)
			if val < minExcludeMaxVal {
				minExcludeMaxVal = val
				row = i + 1
				col = j + 1
			}
		}
	}

	return row, col
}

func getExcludeMaxVal(row int, col int, n int, m int, a [][]int, minExcludeMaxVal int) int {
	maxVal := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i != row && j != col {
				maxVal = max(maxVal, a[i][j])
				if maxVal > minExcludeMaxVal {
					return minExcludeMaxVal
				}
			}
		}
	}
	return maxVal
}
