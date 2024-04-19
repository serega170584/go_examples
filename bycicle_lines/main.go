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
	w, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	h, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	positions := make([][2]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		positions[i][0], _ = strconv.Atoi(scanner.Text())

		scanner.Scan()
		positions[i][1], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(getMinLineSize(w, h, positions))
}

func getMinLineSize(w int, h int, positions [][2]int) int {
	rowPositions := make(map[int][][2]int, h)
	positionsLen := len(positions)
	for i := 0; i < positionsLen; i++ {
		row := positions[i][0]
		rowPositions[row] = append(rowPositions[row], positions[i])
	}

	topRowMinPositions := make([]int, h)
	topRowMaxPositions := make([]int, h)
	topRowMinPosition := math.MaxInt
	topRowMaxPosition := math.MinInt
	for i := 0; i < h; i++ {
		if _, ok := rowPositions[i]; ok {
			curRowPositions := rowPositions[i]
			for _, pos := range curRowPositions {
				topRowMinPosition = min(topRowMinPosition, pos[1])
				topRowMaxPosition = max(topRowMaxPosition, pos[1])
			}
		}
		topRowMinPositions[i] = topRowMinPosition
		topRowMaxPositions[i] = topRowMaxPosition
	}

	bottomRowMinPositions := make([]int, h)
	bottomRowMaxPositions := make([]int, h)
	bottomRowMinPosition := math.MaxInt
	bottomRowMaxPosition := math.MinInt
	for i := h - 1; i >= 0; i-- {
		if _, ok := rowPositions[i]; ok {
			curRowPositions := rowPositions[i]
			for _, pos := range curRowPositions {
				bottomRowMinPosition = min(bottomRowMinPosition, pos[1])
				bottomRowMaxPosition = max(bottomRowMaxPosition, pos[1])
			}
		}
		bottomRowMinPositions[i] = bottomRowMinPosition
		bottomRowMaxPositions[i] = bottomRowMaxPosition
	}

	l := 0
	r := min(w, h)
	for l < r {
		mid := (l + r) / 2
		if check(mid, w-mid, topRowMinPositions, topRowMaxPositions, bottomRowMinPositions, bottomRowMaxPositions) {
			r = mid - 1
		} else {
			l = mid
		}
	}

	return l
}

func check(mid int, maxFirst int, topRowMinPositions []int, topRowMaxPositions []int, bottomRowMinPositions []int, bottomRowMaxPositions []int) bool {
	for i := 0; i <= maxFirst; i++ {
		topRowMinPosition := math.MaxInt
		topRowMaxPosition := math.MinInt
		if i-1 >= 0 {
			topRowMinPosition = topRowMinPositions[i-1]
			topRowMaxPosition = topRowMaxPositions[i-1]
		}

		bottomRowMinPosition := math.MaxInt
		bottomRowMaxPosition := math.MinInt
		next := i + mid
		if next >= 0 {
			bottomRowMinPosition = bottomRowMinPositions[next]
			bottomRowMaxPosition = bottomRowMaxPositions[next]
		}

		minPosition := min(topRowMinPosition, bottomRowMinPosition)
		maxPosition := max(topRowMaxPosition, bottomRowMaxPosition)

		if maxPosition-minPosition < 0 {
			return true
		}

		colLaneWidth := maxPosition - minPosition + 1
		if colLaneWidth <= mid {
			return true
		}
	}

	return false
}
