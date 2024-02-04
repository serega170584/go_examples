package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter cars count")
	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter cars")

	cars := make([][4]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		cars[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		cars[i][1], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		cars[i][2], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		cars[i][3], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Min cars on full parking: ", minCarsOnFullParking(n, 10, cars))
}

func minCarsOnFullParking(n int, maxCars int, cars [][4]int) int {
	minCarsCnt := 11
	nowCars := 0
	occupied := 0
	events := make([][3]int, 2*n)
	for i := 0; i < n; i++ {
		events[2*i] = [3]int{cars[i][0], 1, cars[i][3] - cars[i][2] + 1}
		events[2*i+1] = [3]int{cars[i][1], -1, cars[i][3] - cars[i][2] + 1}
	}

	sort.Slice(events, func(i int, j int) bool {
		return events[i][0] < events[j][0]
	})

	for _, event := range events {
		if event[1] == -1 {
			occupied -= event[2]
			nowCars--
		} else {
			occupied += event[2]
			nowCars++
		}
		if occupied == maxCars {
			minCarsCnt = min(minCarsCnt, nowCars)
		}
	}
	return minCarsCnt
}
