package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	board := make([]string, 8)
	for i := 0; i < 8; i++ {
		scanner.Scan()
		board[i] = scanner.Text()
	}

	fmt.Println(emptySaveCellsCnt(board))
}

func emptySaveCellsCnt(board []string) int {
	bSym := []rune("B")[0]
	rSym := []rune("R")[0]

	runeBoard := make([][]rune, 8)
	for i, row := range board {
		runeRow := []rune(row)
		runeBoard[i] = make([]rune, 8)
		runeBoard[i] = runeRow[0:8]
	}

	filledCells := make([][]bool, 8)
	for i := 0; i < 8; i++ {
		filledCells[i] = make([]bool, 8)
	}

	for i, row := range runeBoard {
		prevRook := -1
		for j := 0; j < 8; j++ {
			v := row[j]

			if prevRook != -1 {
				filledCells[i][j] = true
			}

			if v == bSym {
				filledCells[i][j] = true
				prevRook = -1
			}

			if v == rSym {
				filledCells[i][j] = true
				prevRook = j
			}
		}

		prevRook = -1
		for j := 7; j > -1; j-- {
			v := row[j]

			if prevRook != -1 {
				filledCells[i][j] = true
			}

			if v == bSym {
				filledCells[i][j] = true
				prevRook = -1
			}

			if v == rSym {
				filledCells[i][j] = true
				prevRook = j
			}
		}
	}

	for j := 0; j < 8; j++ {
		prevRook := -1
		for i := 0; i < 8; i++ {
			v := runeBoard[i][j]

			if prevRook != -1 {
				filledCells[i][j] = true
			}

			if v == bSym {
				filledCells[i][j] = true
				prevRook = -1
			}

			if v == rSym {
				filledCells[i][j] = true
				prevRook = j
			}
		}

		prevRook = -1
		for i := 7; i > -1; i-- {
			v := runeBoard[i][j]

			if prevRook != -1 {
				filledCells[i][j] = true
			}

			if v == bSym {
				filledCells[i][j] = true
				prevRook = -1
			}

			if v == rSym {
				filledCells[i][j] = true
				prevRook = j
			}
		}
	}

	for j := 0; j < 8; j++ {
		prevBishop := -1
		for i := 0; i <= 7-j; i++ {
			v := runeBoard[i][i+j]

			if prevBishop != -1 {
				filledCells[i][i+j] = true
			}

			if v == rSym {
				filledCells[i][i+j] = true
				prevBishop = -1
			}

			if v == bSym {
				filledCells[i][i+j] = true
				prevBishop = j
			}
		}

		prevBishop = -1
		for i := 7 - j; i >= 0; i-- {
			v := runeBoard[i][i+j]

			if prevBishop != -1 {
				filledCells[i][i+j] = true
			}

			if v == rSym {
				filledCells[i][i+j] = true
				prevBishop = -1
			}

			if v == bSym {
				filledCells[i][i+j] = true
				prevBishop = j
			}
		}
	}

	for j := 0; j < 8; j++ {
		prevBishop := -1
		for i := j; i <= 7; i++ {
			v := runeBoard[i][i-j]

			if prevBishop != -1 {
				filledCells[i][i-j] = true
			}

			if v == rSym {
				filledCells[i][i-j] = true
				prevBishop = -1
			}

			if v == bSym {
				filledCells[i][i-j] = true
				prevBishop = j
			}
		}

		prevBishop = -1
		for i := 7; i >= j; i-- {
			v := runeBoard[i][i-j]

			if prevBishop != -1 {
				filledCells[i][i-j] = true
			}

			if v == rSym {
				filledCells[i][i-j] = true
				prevBishop = -1
			}

			if v == bSym {
				filledCells[i][i-j] = true
				prevBishop = j
			}
		}
	}

	for j := 0; j < 8; j++ {
		prevBishop := -1
		for i := j; i <= 7; i++ {
			v := runeBoard[i][7-i+j]

			if prevBishop != -1 {
				filledCells[i][7-i+j] = true
			}

			if v == rSym {
				filledCells[i][7-i+j] = true
				prevBishop = -1
			}

			if v == bSym {
				filledCells[i][7-i+j] = true
				prevBishop = 7 - i + j
			}
		}

		prevBishop = -1
		for i := 7; i >= j; i-- {
			v := runeBoard[i][7-i+j]

			if prevBishop != -1 {
				filledCells[i][7-i+j] = true
			}

			if v == rSym {
				filledCells[i][7-i+j] = true
				prevBishop = -1
			}

			if v == bSym {
				filledCells[i][7-i+j] = true
				prevBishop = 7 - i + j
			}
		}
	}

	for j := 0; j < 8; j++ {
		prevBishop := -1
		for i := 0; i <= 7-j; i++ {
			v := runeBoard[i][7-i-j]

			if prevBishop != -1 {
				filledCells[i][7-i-j] = true
			}

			if v == rSym {
				filledCells[i][7-i-j] = true
				prevBishop = -1
			}

			if v == bSym {
				filledCells[i][7-i-j] = true
				prevBishop = 7 - i - j
			}
		}

		prevBishop = -1
		for i := 7 - j; i >= 0; i-- {
			v := runeBoard[i][7-i-j]

			if prevBishop != -1 {
				filledCells[i][7-i-j] = true
			}

			if v == rSym {
				filledCells[i][7-i-j] = true
				prevBishop = -1
			}

			if v == bSym {
				filledCells[i][7-i-j] = true
				prevBishop = 7 - i - j
			}
		}
	}

	emptyCellsCnt := 0

	for _, row := range filledCells {
		for _, v := range row {
			if !v {
				emptyCellsCnt++
			}
		}
	}

	return emptyCellsCnt
}
