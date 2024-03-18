package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//10
//4 4
//10 2
//5 5
//5 1
//1 8
//9 3
//9 6
//8 5
//1 9
//4 5

// 0 0 0 0 0 0 0 x x 0
// 0 0 0 0 0 0 0 0 0 0
// 0 0 0 0 0 0 0 0 0 0
// 0 0 0 x x 0 0 0 0 0
// x 0 0 0 x 0 0 0 0 0
// 0 0 0 0 0 0 0 0 0 0
// 0 0 0 0 0 0 0 0 0 0
// 0 0 0 0 x 0 0 0 0 0
// 0 0 x 0 0 x 0 0 0 0
// 0 x 0 0 0 0 0 0 0 0
// 3 + 5 + 2 + 5 + 4 + 1 + 3 = 23

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
	colMaxVals := make([]int, m)
	colMax2Vals := make([]int, m)
	colMaxIndexes := make([]int, m)
	for j := 0; j < m; j++ {
		if a[1][j] > a[0][j] {
			colMaxIndexes[j] = 1
			colMaxVals[j] = a[1][j]
			colMax2Vals[j] = a[0][j]
		} else {
			colMaxIndexes[j] = 0
			colMaxVals[j] = a[0][j]
			colMax2Vals[j] = a[1][j]
		}
	}

	for i := 2; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] > colMaxVals[j] {
				colMaxIndexes[j] = i
				colMax2Vals[j] = colMaxVals[j]
				colMaxVals[j] = a[i][j]
			} else if a[i][j] > colMax2Vals[j] {
				colMax2Vals[j] = a[i][j]
			}
		}
	}

	maxRowVal := 0
	max2RowVal := 0
	row := -1
	col := -1
	for j := 0; j < m; j++ {
		curMaxVal := colMaxVals[j]
		if colMaxIndexes[j] == 0 {
			curMaxVal = colMax2Vals[j]
		}
		if curMaxVal > maxRowVal {
			max2RowVal = maxRowVal
			maxRowVal = curMaxVal
			row = 1
			col = j + 1
		} else if curMaxVal > max2RowVal {
			max2RowVal = curMaxVal
		}
	}

	for i := 1; i < n; i++ {
		baseMax2RowVal := 0
		baseMaxRowVal := 0
		baseRow := -1
		baseCol := -1
		for j := 0; j < m; j++ {
			curMaxVal := colMaxVals[j]
			if colMaxIndexes[j] == i {
				curMaxVal = colMax2Vals[j]
			}

			if curMaxVal > baseMaxRowVal {
				baseMax2RowVal = baseMaxRowVal
				baseMaxRowVal = curMaxVal
				baseRow = i + 1
				baseCol = j + 1
			} else if curMaxVal > baseMax2RowVal {
				baseMax2RowVal = curMaxVal
			}

			if baseMax2RowVal >= max2RowVal {
				break
			}
		}

		if baseMax2RowVal < max2RowVal {
			row = baseRow
			col = baseCol
			max2RowVal = baseMax2RowVal
			maxRowVal = baseMaxRowVal
		}
	}

	return row, col
}
