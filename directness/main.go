package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 3 4
// 5 6 5 - 3 = 2 6 - 4 = 2
// kx + b, 3k + b = 4
// 5k + b = 6,
// 5k + 4 - 3k = 6
// k + 2 = 3
// k = 1, b = 1

// (3k + b)/5k + b , 3k / 5k + b + b / 5k + b = 3k + b + 2k - 2k

// kx = y - b
// (y - b) / k = x,  y / k - b / k

// 2/3
// 40 60 20
// 60 90 30

// define start x, start y (point with min x)
// each point - startx, - starty
// fix point with start x 0, if start y !=0 each y - start y
// create array of segments: x = point x - prev point x, y = point y - prev point y
// watch first segment: least, most coord: x or y, all segment should have same kind of least of most coord
// calculate max shared delimiter and use it to next coordinate
// if delimiter == 1 or least > most then break loop and print NO
// after pass all the cycle print YES
func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Scan()
	cnt, _ := strconv.Atoi(s.Text())

	points := make([][]int, cnt)
	for i := 0; i < cnt; i++ {
		points[i] = make([]int, 2)
		s.Scan()
		points[i][0], _ = strconv.Atoi(s.Text())
		s.Scan()
		points[i][1], _ = strconv.Atoi(s.Text())
	}

	minXPoint := make([]int, 2)
	minXPoint[0] = 1000000000001
	for _, point := range points {
		if point[0] < minXPoint[0] {
			minXPoint[0] = point[0]
			minXPoint[1] = point[1]
		}
	}

	for i := range points {
		points[i][0] -= minXPoint[0]
		points[i][1] -= minXPoint[1]
		if points[i][1] < 0 {
			points[i][1] = -points[i][1]
		}
	}

	fixPoint := make([]int, 2)
	fixPoint[0] = 1000000000001
	fixPoint[1] = 1000000000001
	isNo := false

	for _, point := range points {
		if point[0] == 0 && point[1] == 0 {
			continue
		}

		if fixPoint[0] == 0 && point[0] != 0 {
			isNo = true
			fmt.Println("NO")
			break
		}

		if fixPoint[1] == 0 && point[1] != 0 {
			isNo = true
			fmt.Println("NO")
			break
		}

		if fixPoint[0] == 0 || fixPoint[1] == 0 {
			continue
		}

		if fixPoint[0] == 1000000000001 {
			fixPoint[0] = point[0]
		}

		if fixPoint[1] == 1000000000001 {
			fixPoint[1] = point[1]
		}

		if fixPoint[0] == point[0] && fixPoint[1] == point[1] {
			continue
		}

		if point[1] == point[0] && fixPoint[1] == fixPoint[0] {
			continue
		}

		sharedDelimiterX := maxSharedDelimiter(point[0], fixPoint[0])
		sharedDelimiterY := maxSharedDelimiter(point[1], fixPoint[1])

		if point[0]/sharedDelimiterX != point[1]/sharedDelimiterY {
			isNo = true
			fmt.Println("NO")
			break
		}

		if fixPoint[0]/sharedDelimiterX != fixPoint[1]/sharedDelimiterY {
			isNo = true
			fmt.Println("NO")
			break
		}
	}

	if !isNo {
		fmt.Println("YES")
	}

}

func maxSharedDelimiter(num int, delimiter int) int {
	if num == delimiter {
		return delimiter
	}

	if delimiter > num {
		delimiter, num = num, delimiter
	}

	for delimiter != 0 && 1 != delimiter {
		num, delimiter = delimiter, num%delimiter
	}

	if delimiter == 0 {
		delimiter = num
	}

	return delimiter
}
