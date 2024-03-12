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

	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(leftMinRopeLength(list))
}

func leftMinRopeLength(list []int) int {
	maxVal := math.MinInt
	sum := 0
	for _, v := range list {
		maxVal = max(maxVal, v)
		sum += v
	}

	maxQuantity := 0
	for _, v := range list {
		if maxVal == v {
			maxQuantity++
		}
	}

	if maxQuantity == 1 {
		leftRopeLength := maxVal
		for _, v := range list {
			if v != maxVal {
				leftRopeLength -= v
			}
		}
		if leftRopeLength > 0 {
			return leftRopeLength
		}
	}
	return sum
}
