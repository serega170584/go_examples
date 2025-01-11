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

	minTop := math.MaxInt
	maxTop := math.MinInt
	minLeft := math.MaxInt
	maxLeft := math.MinInt
	for _, v := range list {
		minTop = min(minTop, v[0])
		maxTop = max(maxTop, v[0])
		minLeft = min(minLeft, v[1])
		maxLeft = max(maxLeft, v[1])
	}

	rect := make([][]bool, 0, maxTop-minTop+1)
	for i := 0; i < maxTop-minTop+1; i++ {
		rect = append(rect, make([]bool, maxLeft-minLeft+1))
	}

	for _, v := range list {
		rect[v[0]-minTop][v[1]-minLeft] = true
	}

	perimeter := 0
	for i, row := range rect {
		for j, v := range row {
			if j == 0 && v {
				perimeter++
			}

			if j == maxLeft-minLeft && v {
				perimeter++
			}

			if i == 0 && v {
				perimeter++
			}

			if i == maxTop-minTop && v {
				perimeter++
			}

			if j != 0 && !rect[i][j-1] && v {
				perimeter++
			}

			if j != maxLeft-minLeft && !rect[i][j+1] && v {
				perimeter++
			}

			if i != 0 && !rect[i-1][j] && v {
				perimeter++
			}

			if i != maxTop-minTop && !rect[i+1][j] && v {
				perimeter++
			}
		}
	}

	fmt.Println(perimeter)
}
