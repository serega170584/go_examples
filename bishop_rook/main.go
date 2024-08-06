package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cells := [8][8]string{}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < 8; i++ {
		scanner.Scan()
		t := scanner.Text()
		for j, v := range t {
			cells[i][j] = string([]rune{v})
		}
	}

	fmt.Println(getEmptyCellsCnt(cells))
}

func getEmptyCellsCnt(cells [8][8]string) int {
	ltBishops := [8][8]bool{}
	rtBishops := [8][8]bool{}
	lbBishops := [8][8]bool{}
	rbBishops := [8][8]bool{}

	lRooks := [8][8]bool{}
	rRooks := [8][8]bool{}
	tRooks := [8][8]bool{}
	bRooks := [8][8]bool{}

	cnt := 0

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if cells[i][j] == "R" {
				lRooks[i][j] = true
				tRooks[i][j] = true
				ltBishops[i][j] = false
				rtBishops[i][j] = false
				continue
			}

			if cells[i][j] == "B" {
				ltBishops[i][j] = true
				rtBishops[i][j] = true
				lRooks[i][j] = false
				tRooks[i][j] = false
				continue
			}

			if j >= 1 && lRooks[i][j-1] {
				lRooks[i][j] = true
			}

			if i >= 1 && tRooks[i-1][j] {
				tRooks[i][j] = true
			}

			if i >= 1 && j >= 1 && ltBishops[i-1][j-1] {
				ltBishops[i][j] = true
			}

			if i >= 1 && j <= 6 && rtBishops[i-1][j+1] {
				rtBishops[i][j] = true
			}
		}
	}

	for i := 7; i >= 0; i-- {
		for j := 7; j >= 0; j-- {
			if cells[i][j] == "R" {
				rRooks[i][j] = true
				bRooks[i][j] = true
				lbBishops[i][j] = false
				rbBishops[i][j] = false
				continue
			}

			if cells[i][j] == "B" {
				lbBishops[i][j] = true
				rbBishops[i][j] = true
				rRooks[i][j] = false
				bRooks[i][j] = false
				continue
			}

			if j <= 6 && rRooks[i][j+1] {
				rRooks[i][j] = true
			}

			if i <= 6 && bRooks[i+1][j] {
				bRooks[i][j] = true
			}

			if i <= 6 && j >= 1 && lbBishops[i+1][j-1] {
				lbBishops[i][j] = true
			}

			if i <= 6 && j <= 6 && rbBishops[i+1][j+1] {
				rbBishops[i][j] = true
			}
		}
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if ltBishops[i][j] {
				continue
			}

			if rtBishops[i][j] {
				continue
			}

			if lbBishops[i][j] {
				continue
			}

			if rbBishops[i][j] {
				continue
			}

			if tRooks[i][j] {
				continue
			}

			if lRooks[i][j] {
				continue
			}

			if bRooks[i][j] {
				continue
			}

			if rRooks[i][j] {
				continue
			}

			cnt++
		}
	}

	return cnt
}
