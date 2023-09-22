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

// 1 1 1 1 1 1
// 1 1 1 1 1 1
// 1 1 1 1 1 1
// 1 1 1 1 1 1
// 1 1 1 1 1 1
// 1 1 1 1 1 1

// 1 1 1 1
// 1 1 1 1
// 1 1 1 1
// 1 0 0 1

// [1] [1 2] [1 3] [1 4]
// [1 2 3] [1 3 4 2]

// 1
// [1] [1 2] [1 3] [1 4] [1 5] [1 6]

// 2
// [2] [2]

// position
// positions
// cnt

// 0 0 1 0
// 1 1 1 1
// 1 1 1 1
// 1 0 1 0

// 0 0 0 1
// 0 0 1 1
// 1 1 1 1
// 1 1 0 1

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

	for i := 0; i < cnt; i++ {
		copyDispositions := make([]int, cnt)
		copy(copyDispositions, dispositions)
		copyDispositions[0] = i + 1

		copyBusyPositions := make([][]bool, cnt)
		for j := 1; j < cnt; j++ {
			copyBusyPositions[j] = make([]bool, cnt)
			copy(copyBusyPositions[j], busyPositions[j])

			copyBusyPositions[j][i] = true
			if i-j > -1 {
				copyBusyPositions[j][i-j] = true
			}
			if i+j < cnt {
				copyBusyPositions[j][i+j] = true
			}
		}

		generateDisposition(copyBusyPositions, 1, cnt, copyDispositions)
	}
}

func generateDisposition(busyPositions [][]bool, rowInd, cnt int, dispositions []int) {
	for i := 0; i < cnt; i++ {
		if busyPositions[rowInd][i] {
			continue
		}

		copyDispositions := make([]int, cnt)
		copy(copyDispositions, dispositions)
		copyDispositions[rowInd] = i + 1

		if rowInd == cnt-1 {
			str := make([]string, cnt)
			for dispInd, val := range copyDispositions {
				str[dispInd] = strconv.Itoa(val)
			}
			fmt.Println(strings.Join(str, " "))
			continue
		}

		copyBusyPositions := make([][]bool, cnt)
		for j := rowInd + 1; j < cnt; j++ {
			copyBusyPositions[j] = make([]bool, cnt)
			copy(copyBusyPositions[j], busyPositions[j])

			copyBusyPositions[j][i] = true
			if i-j+rowInd > -1 {
				copyBusyPositions[j][i-j+rowInd] = true
			}
			if i+j-rowInd < cnt {
				copyBusyPositions[j][i+j-rowInd] = true
			}
		}

		generateDisposition(copyBusyPositions, rowInd+1, cnt, copyDispositions)
	}
}
