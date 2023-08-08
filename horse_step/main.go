package main

import (
	"fmt"
	"log"
)

func main() {
	var rowsCnt, colsCnt int

	_, err := fmt.Scanln(&rowsCnt, &colsCnt)
	if err != nil {
		log.Fatal(err)
	}

	board := make([][]int, rowsCnt)

	for i := 0; i < rowsCnt; i++ {
		board[i] = make([]int, colsCnt)
	}

	board[0][0] = 1

	for i := 1; i < rowsCnt; i++ {
		for j := 1; j < colsCnt; j++ {
			if i == 1 && j == 1 {
				continue
			}

			leftPathCnt := 0
			if j > 1 {
				leftPathCnt = board[i-1][j-2]
			}

			upPathCnt := 0
			if i > 1 {
				upPathCnt = board[i-2][j-1]
			}

			board[i][j] = leftPathCnt + upPathCnt
		}
	}

	fmt.Println(board[rowsCnt-1][colsCnt-1])
}
