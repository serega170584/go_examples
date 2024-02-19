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

	cost, path := maxCost(n, list, stepsCnt)
	fmt.Println("Got max cost", cost)

	fmt.Println("Got path", path)
}

func maxCost(n int, list []int, stepsCnt int) (int, []int) {
	cost := make([]int, n)
	cost[0] = list[0]
	path := make([]int, n)
	for i := 1; i < n; i++ {
		maxVal := -math.MaxInt
		for j := 0; j < stepsCnt; j++ {
			prev := i - j - 1
			if prev >= 0 {
				if cost[prev]+list[i] > maxVal {
					maxVal = cost[prev] + list[i]
					path[i] = prev
				}
			}
		}
		cost[i] = maxVal
	}

	fmt.Println(cost)
	fmt.Println(path)

	prevPath := make([]int, n)
	ind := n - 1
	cnt := 0
	for ind != 0 {
		prevPath[ind] = path[ind]
		ind = path[ind]
		cnt++
	}

	resultPath := make([]int, cnt-1)
	ind = 0
	for _, v := range prevPath {
		if v != 0 {
			resultPath[ind] = v
			ind++
		}
	}

	return cost[n-1], resultPath
}
