package main

import (
	"bufio"
	"fmt"
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

	a := make([][]int, m)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			scanner.Scan()
			a[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}

	fmt.Println(getRectangleArea(a))
}

func getRectangleArea(A [][]int) int {
	hist := make([][]int, len(A))
	hist[0] = make([]int, len(A[0]))
	for i := 0; i < len(A[0]); i++ {
		hist[0][i] = A[0][i]
	}

	for i := 1; i < len(A); i++ {
		hist[i] = make([]int, len(A[i]))
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] == 1 {
				hist[i][j] = hist[i-1][j] + 1
			}
		}
	}

	maxArea := 0
	for i := 0; i < len(hist); i++ {
		cnt := make([]int, len(hist)+1)
		for j := 0; j < len(hist[i]); j++ {
			cnt[hist[i][j]]++
		}

		currInd := 0
		for k := len(hist); k >= 0; k-- {
			if cnt[k] != 0 {
				for l := 0; l < cnt[k]; l++ {
					hist[i][currInd] = k
					currInd++
				}
			}
		}

		for k := currInd; k < len(A[i]); k++ {
			hist[i][k] = 0
		}
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if hist[i][j]*(j+1) > maxArea {
				maxArea = hist[i][j] * (j + 1)
			}
		}
	}

	return maxArea
}
