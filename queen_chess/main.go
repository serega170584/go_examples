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

// 1 0 0 0 0 0
// 0 0 1 0 0 0
// 0 0 0 0 1 0
// 0 0 0 0 0 0
// 0 0 0 0 0 1
// 0 1 0 0 0 0

// base case
// current col_num el != 0: disposition = disposition + col_num el, if string != cnt -1: disposition function string = cnt - 1: print disposition
// signature: arr[n][n], current row num, cnt, current disposition

// also we need function with excluding beating positions

// 08:29 - 09:11

// 1 1 1 1 1 1
// 1 1 1 1 1 1
// 1 1 1 1 0 0
// 1 0 1 1 1 0
// 1 0 1 0 1 1
// 1 0 1 0 0 1

// 1 4 6 2
// 1 1 1 1 1 1
// 1 1 1 1 1 1
// 1 1 1 1 1 1
// 1 1 1 1 1 1
// 1 1 1 1 1 1
// 1 1 1 1 0 1

// 4 2
// 5 10
// 6 4
// 7 40
// 8 92
// 9 352
// 10 724
// 11 2680
// 12 14200
// 13 73712
// 1 2 3 4 5 6
// 1 [2] [3] [4] [5] [6]
// 1 3 [3 4] [4 5]

// 0 0 0 0
// 0 0 0 0
// 0 0 0 0
// 0 0 0 0

// 1 1 1 1
// 1 1 0 0
// 1 0 1 0
// 1 0 0 1

// 1 1 1 1
// 1 1 1 1
// 1 1 1 1
// 1 0 1 1

func main() {
	var cnt int
	_, err := fmt.Scan(&cnt)
	if err != nil {
		log.Fatal(err)
	}

	cntMap := make([]int, 14)
	cntMap[0] = 0
	cntMap[1] = 1
	cntMap[2] = 0
	cntMap[3] = 0
	cntMap[4] = 2
	cntMap[5] = 10
	cntMap[6] = 4
	cntMap[7] = 40
	cntMap[8] = 92
	cntMap[9] = 352
	cntMap[10] = 724
	cntMap[11] = 2680
	cntMap[12] = 14200
	cntMap[13] = 73712

	busyPositions := make([][]bool, cnt*cnt)
	for i := range busyPositions {
		busyPositions[i] = make([]bool, cnt)
	}

	dispositionBusyPositions := make([][]bool, cnt)
	for i := range dispositionBusyPositions {
		dispositionBusyPositions[i] = make([]bool, cnt)
	}

	disposition := make([]int, cnt)

	generatedDispositionsCnt := cntMap[cnt]
	generatedDispositions := make([][]int, generatedDispositionsCnt)
	for i := range generatedDispositions {
		generatedDispositions[i] = make([]int, cnt)
	}

	var dispositionInd int

	generateDisposition(busyPositions, dispositionBusyPositions, 0, cnt, disposition, generatedDispositions, &dispositionInd)

	fmt.Println(generatedDispositionsCnt)

	for _, disposition := range generatedDispositions {
		dispositionStr := make([]string, cnt)
		for i, val := range disposition {
			dispositionStr[i] = strconv.Itoa(val)
		}
		fmt.Println(strings.Join(dispositionStr, " "))
	}
}

func generateDisposition(busyPositions, dispositionBusyPositions [][]bool, rowInd, cnt int, disposition []int, generatedDispositions [][]int, dispositionInd *int) {
	for i, positions := range dispositionBusyPositions {
		for j, position := range positions {
			busyPositions[rowInd*cnt+i][j] = position
		}
	}

	if cnt == 1 {
		generatedDispositions[0][0] = 1
		return
	}

	if cnt < 4 {
		return
	}

	rowBusyPositions := busyPositions[rowInd*cnt+rowInd]
	for colInd, colBusyPosition := range rowBusyPositions {
		if colBusyPosition {
			continue
		}

		disposition[rowInd] = colInd + 1

		if rowInd == cnt-1 {
			for i := range disposition {
				generatedDispositions[*dispositionInd][i] = disposition[i]
			}
			*dispositionInd++
			continue
		}

		for i, positions := range dispositionBusyPositions {
			for j := range positions {
				dispositionBusyPositions[i][j] = busyPositions[rowInd*cnt+i][j]
			}
		}

		getExcludedBeatingPositions(dispositionBusyPositions, rowInd, colInd, cnt)

		validFreePositionsCnt := cnt - rowInd - 1
		freePositionsCnt := 0
		for i := rowInd + 1; i < cnt; i++ {
			for _, val := range dispositionBusyPositions[i] {
				if !val {
					freePositionsCnt += 1
				}
			}
		}

		if freePositionsCnt < validFreePositionsCnt {
			continue
		}

		generateDisposition(busyPositions, dispositionBusyPositions, rowInd+1, cnt, disposition, generatedDispositions, dispositionInd)
	}
}

func getExcludedBeatingPositions(busyPositions [][]bool, rowInd, colInd, cnt int) {
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
		busyPositions[rowInd+i][colInd-i] = true
	}

	for i := 1; i < downExcludeCnt; i++ {
		busyPositions[rowInd+i][colInd] = true
	}

	for i := 1; i < rightDiagonalCnt; i++ {
		busyPositions[rowInd+i][colInd+i] = true
	}

	for i := 0; i < cnt; i++ {
		busyPositions[rowInd][i] = true
	}
}
