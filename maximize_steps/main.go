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

	fmt.Println("Enter list count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list")
	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Enter steps count")
	scanner.Scan()
	stepsCnt, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Got max cost", maxCost(n, list, stepsCnt))
}

func maxCost(n int, list []int, stepsCnt int) int {
	cost := make([]int, n)
	cost[0] = list[0]
	for i := 1; i < n; i++ {
		maxVal := -math.MaxInt
		for j := 0; j < stepsCnt; j++ {
			prev := i - j - 1
			if prev >= 0 {
				maxVal = max(cost[prev]+list[i], maxVal)
			}
		}
		cost[i] = maxVal
	}
	fmt.Println(cost)
	return cost[n-1]
}
