package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// 0 0 0 x 0 0
// 0 x 0 0 0 0
// 0 0 0 0 0 x
// 0 0 0 0 0 x
// 0 0 0 0 0 x
// 0 0 x 0 0 0

// x 0 0 x 0 0
// 0 x 0 0 0 0
// 0 0 0 0 x x
// 0 0 0 0 0 x
// 0 0 0 0 0 0
// 0 0 0 0 0 0
// (4 - 0) + (5 - 2)
// (4 - 2) + (5 - 0)

// x 0 0 x 0 0
// 0 x 0 0 0 0
// 0 0 0 0 0 0
// 0 0 0 0 0 x
// 0 0 0 0 x x
// 0 0 0 0 0 0
// (2 - 0) + (5 - 4)

// 0 0 0 0 0 0
// 0 x 0 0 0 0
// x 0 0 x 0 0
// 0 0 0 0 0 x
// 0 0 0 0 x x
// 0 0 0 0 0 0
// (2 - 0) + (5 - 4)

// 0 0 0 0 x
// 0 0 0 x 0
// 0 0 x 0 0
// 0 x 0 0 0
// x 0 0 0 0

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	a := make([][2]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		a[i][1], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(getMinSteps(n, a))
}

func getMinSteps(n int, a [][2]int) int {
	cornerLeft := math.MaxInt
	cornerRight := math.MinInt

	for _, v := range a {
		cornerLeft = min(v[1], cornerLeft)
		cornerRight = max(v[1], cornerRight)
	}

	minHorizontalDist := math.MaxInt
	for i := cornerLeft; i <= cornerRight; i++ {
		horizontalSteps := 0
		for _, v := range a {
			dist := v[1] - i
			if dist < 0 {
				dist = -dist
			}
			horizontalSteps += dist

			if horizontalSteps > minHorizontalDist {
				break
			}
		}
		minHorizontalDist = min(minHorizontalDist, horizontalSteps)
	}

	cnts := make([]int, n)
	for _, v := range a {
		cnts[v[0]-1]++
	}

	extraFilled := make([]int, 0)
	zero := make([]int, 0)
	specialCnt := 0
	for i, v := range cnts {
		if v == 0 {
			zero = append(zero, i)
			specialCnt++
		} else if v > 1 {
			for j := 0; j < v-1; j++ {
				extraFilled = append(extraFilled, i)
				specialCnt++
			}
		}
	}

	specialCnt /= 2

	minVerticalDist := 0
	for i := 0; i < specialCnt; i++ {
		dist := extraFilled[i] - zero[i]
		if dist < 0 {
			dist = -dist
		}
		minVerticalDist += dist
	}

	return minVerticalDist + minHorizontalDist
}
