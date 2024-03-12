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
	pairs := make([][2]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		pairs[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		pairs[i][1], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(getMinRectanglePositions(pairs))
}

func getMinRectanglePositions(pairs [][2]int) (int, int, int, int) {
	minXVal := math.MaxInt
	maxXVal := math.MinInt
	minYVal := math.MaxInt
	maxYVal := math.MinInt
	for _, v := range pairs {
		minXVal = min(minXVal, v[0])
		maxXVal = max(maxXVal, v[0])
		minYVal = min(minYVal, v[1])
		maxYVal = max(maxYVal, v[1])
	}
	return minXVal, minYVal, maxXVal, maxYVal
}
