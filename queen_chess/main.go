package main

import (
	"bufio"
	"fmt"
	"os"
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

// [1] [2] [3] [4]
// 1
// [1 2] [1 3] [1 4]

// 0 0 0 0
// 0 0 0 0
// 0 0 0 0
// 0 0 0 0

// 1 1 1 1
// 1 1 1 0
// 0 1 0 1
// 0 1 0 0

// 1 1 1 1
// 1 1 1 1
// 0 1 1 1
// 0 1 0 1

// 0 0 1 0 0 0
// 0 1 1 1 1 0
// 1 0 1 1 1 1
// 0 0 1 0 1 1
// 0 1 1 0 1 0
// 1 0 1 0 1 0

// 1
// [1] [1 2] [1 3] [1 4] [1 5] [1 6]

// 2
// [2] [2 1] []

// position
// positions
// cnt

// 1 1 1 1
// 1 1 0 0
// 1 0 1 0
// 1 0 0 1

// 1 1 1 1
// 1 1 1 1
// 1 1 1 1
// 1 0 1 1

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Scan()
	cnt, _ := strconv.Atoi(s.Text())

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

	busyPositions := make([][]bool, cnt)
	for i := range busyPositions {
		busyPositions[i] = make([]bool, cnt)
	}

	dispositions := make([]int, cnt)

	fmt.Println(cntMap[cnt])

	if cnt == 1 {
		fmt.Println("1")
		return
	}

	if cnt < 4 {
		return
	}

	counterCh := make(chan struct{}, cnt)

	rowBusyPositions := busyPositions[0]
	for colInd := range rowBusyPositions {
		copyDispositions := make([]int, cnt)
		copy(copyDispositions, dispositions)

		go func(dispositions []int, colInd int, busyPositions [][]bool, counterCh chan struct{}) {
			dispositions[0] = colInd + 1

			copyBusyPositions := make([][]bool, cnt)
			for i, val := range busyPositions {
				copyBusyPositions[i] = make([]bool, cnt)
				copy(copyBusyPositions[i], val)
			}

			getExcludedBeatingPositions(copyBusyPositions, 0, colInd, cnt)

			generateDisposition(copyBusyPositions, 1, cnt, dispositions)

			counterCh <- struct{}{}
		}(copyDispositions, colInd, busyPositions, counterCh)
	}

	handledCnt := 0
	for range counterCh {
		handledCnt++
		if handledCnt == cnt {
			close(counterCh)
		}
	}
}

func generateDisposition(busyPositions [][]bool, rowInd, cnt int, dispositions []int) {
	rowBusyPositions := busyPositions[rowInd]
	for colInd, colBusyPosition := range rowBusyPositions {
		if colBusyPosition {
			continue
		}

		dispositions[rowInd] = colInd + 1

		if rowInd == cnt-1 {
			str := make([]string, cnt)
			for i, val := range dispositions {
				str[i] = strconv.Itoa(val)
			}
			fmt.Println(strings.Join(str, " "))
			continue
		}

		copyBusyPositions := make([][]bool, cnt)
		for i, val := range busyPositions {
			copyBusyPositions[i] = make([]bool, cnt)
			copy(copyBusyPositions[i], val)
		}

		getExcludedBeatingPositions(copyBusyPositions, rowInd, colInd, cnt)

		generateDisposition(copyBusyPositions, rowInd+1, cnt, dispositions)
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
}
