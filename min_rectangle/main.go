package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	points := [][2]int{}

	for i := 0; i < n; i++ {
		points = append(points, [2]int{})

		scanner.Scan()
		points[i][0], _ = strconv.Atoi(scanner.Text())

		scanner.Scan()
		points[i][1], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(getRect(points))
}

func getRect(points [][2]int) string {
	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt
	for _, v := range points {
		minX = min(minX, v[0])
		maxX = max(maxX, v[0])
		minY = min(minY, v[1])
		maxY = max(maxY, v[1])
	}

	res := []string{strconv.Itoa(minX), strconv.Itoa(minY), strconv.Itoa(maxX), strconv.Itoa(maxY)}
	return strings.Join(res, " ")
}
