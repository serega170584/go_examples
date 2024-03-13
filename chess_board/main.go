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
	n, _ := strconv.Atoi(scanner.Text())

	list := make([][2]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		list[i][1], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(getPerimeter(list))
}

func getPerimeter(list [][2]int) int {
	minXVal := math.MaxInt
	maxXVal := math.MinInt
	minYVal := math.MaxInt
	maxYVal := math.MinInt
	for _, v := range list {
		minXVal = min(minXVal, v[0])
		maxXVal = max(maxXVal, v[0])
		minYVal = min(minYVal, v[1])
		maxYVal = max(maxYVal, v[1])
	}

	width := maxXVal - minXVal + 1
	height := maxYVal - minYVal + 1

	rect := make([][]bool, height)
	for i := 0; i < height; i++ {
		rect[i] = make([]bool, width)
	}

	for _, v := range list {
		rect[v[1]-minYVal][v[0]-minXVal] = true
	}

	perimeter := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 && rect[i][j] {
				perimeter++
			}

			if j == 0 && rect[i][j] {
				perimeter++
			}

			if i+1 == height && rect[i][j] {
				perimeter++
			}

			if j+1 == width && rect[i][j] {
				perimeter++
			}

			if i != height-1 && !rect[i+1][j] && rect[i][j] {
				perimeter++
			}

			if j != width-1 && !rect[i][j+1] && rect[i][j] {
				perimeter++
			}

			if i != 0 && !rect[i-1][j] && rect[i][j] {
				perimeter++
			}

			if j != 0 && !rect[i][j-1] && rect[i][j] {
				perimeter++
			}
		}
	}

	return perimeter
}
