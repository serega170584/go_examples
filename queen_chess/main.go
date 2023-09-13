package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// 2 4 1 3
// 1 0 1 1
// 1 1 1 0
// 0 1 1 1
// 1 1 0 1

// 3 1 4 2
// 1 1 0 1
// 0 1 1 1
// 1 1 1 0
// 1 0 1 1

// 0 1 1 1
// 1 1 1 0
// 1 0 1 1
// 1 1 1 1

// 1 1 - 2 3, 3 2, 2 4, 4 2
// 2 1 - 3 3, 4 2, 4 3
// 3 1 - 2 3, 4 3, 4 4, 1 2, 1 4,

// 3 1
// 0 1 0 1
// 0 0 1 1
// 0 0 0 0
// 0 0 1 1

// 1 2, 2 1, 1 3, 3 1, 2 3, 3 2
// -2 +1,

// 1 1 1 1
// 1 1 1 1
// 1 1 1 1
// 1 1 1 1

// 0 0 0 0 - 1 2
// 0 0 0 1
// 1 0 1 0
// 1 0 1 1

// 0 0 0 0 - 1 2
// 0 0 0 0 - 2 4
// 1 0 0 0
// 1 0 1 0

// 0 0 0 0 - 1 2
// 0 0 0 0 - 2 4
// 0 0 0 0 - 3 1
// 0 0 1 0 - 4 3

// 0 0 0 0 1
// 0 0 1 0 0
// 1 0 0 0 0
// 0 0 0 1 0
// 0 1 0 0 0

// base case
// current col_num el != 0: disposition = disposition + col_num el, if string != cnt -1: disposition function string = cnt - 1: print disposition
// signature: arr[n][n], current row num, cnt, current disposition

// also we need function with excluding beating positions

// 08:29 - 09:11

func main() {
	var cnt int
	_, err := fmt.Scan(&cnt)
	if err != nil {
		log.Fatal(err)
	}

	busyPositions := make([][]bool, cnt)
	for i := range busyPositions {
		busyPositions[i] = make([]bool, cnt)
	}

	disposition := make([]int, cnt)

	generateDisposition(busyPositions, 0, cnt, disposition)
}

func generateDisposition(busyPositions [][]bool, rowInd, cnt int, disposition []int) {
	if cnt < 4 {
		return
	}

	rowBusyPositions := busyPositions[rowInd]
	for colInd, colBusyPosition := range rowBusyPositions {
		if colBusyPosition {
			continue
		}

		localDisposition := make([]int, cnt)
		copy(localDisposition, disposition)
		localDisposition[rowInd] = colInd + 1

		if rowInd == cnt-1 {
			stringDisposition := make([]string, cnt)
			for i, val := range localDisposition {
				stringDisposition[i] = strconv.Itoa(val)
			}
			fmt.Println(strings.Join(stringDisposition, " "))
			return
		}

		excludedBeatingPositions := getExcludedBeatingPositions(busyPositions, rowInd, colInd, cnt)

		generateDisposition(excludedBeatingPositions, rowInd+1, cnt, localDisposition)

	}
}

func getExcludedBeatingPositions(busyPositions [][]bool, rowInd, colInd, cnt int) [][]bool {
	localBusyPositions := make([][]bool, cnt)
	for i := range localBusyPositions {
		localBusyPositions[i] = make([]bool, cnt)
		copy(localBusyPositions[i], busyPositions[i])
	}

	leftExcludeCnt := colInd + 1
	downExcludeCnt := cnt - rowInd
	rightExcludeCnt := cnt - colInd

	leftDiagonalCnt := leftExcludeCnt
	if leftExcludeCnt > downExcludeCnt {
		leftDiagonalCnt = downExcludeCnt
	}

	rightDiagonalCnt := rightExcludeCnt
	if rightExcludeCnt > downExcludeCnt {
		rightDiagonalCnt = downExcludeCnt
	}

	for i := 1; i < leftDiagonalCnt; i++ {
		localBusyPositions[rowInd+i][colInd-i] = true
	}

	for i := 1; i < downExcludeCnt; i++ {
		localBusyPositions[rowInd+i][colInd] = true
	}

	for i := 1; i < rightDiagonalCnt; i++ {
		localBusyPositions[rowInd+i][colInd+i] = true
	}

	for i := 0; i < cnt; i++ {
		localBusyPositions[rowInd][i] = true
	}

	return localBusyPositions
}
