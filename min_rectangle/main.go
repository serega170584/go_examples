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
	k, _ := strconv.Atoi(scanner.Text())

	list := make([][2]int, k)
	for i := 0; i < k; i++ {
		scanner.Scan()
		list[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		list[i][1], _ = strconv.Atoi(scanner.Text())
	}

	minLeft := math.MaxInt
	maxLeft := math.MinInt
	minTop := math.MaxInt
	maxTop := math.MinInt

	for _, v := range list {
		minLeft = min(minLeft, v[0])
		maxLeft = max(maxLeft, v[0])
		minTop = min(minTop, v[1])
		maxTop = max(maxTop, v[1])
	}

	fmt.Println(minLeft, minTop, maxLeft, maxTop)
}
