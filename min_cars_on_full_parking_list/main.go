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
	cars := make([][5]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		cars[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		cars[i][1], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		cars[i][2], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		cars[i][3], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		cars[i][4], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Get min cars", minCarsOnFullParkingList(n, 10, cars))
}

func minCarsOnFullParkingList(n int, maxCars int, cars [][5]int) []int {
	events := make([][4]int, 2*n)
	minCarsList := make([]int, n)
	occupied := 0
	minCarsCnt := 11
	minCarIndex := 0
	nowCars := 0

	for i := 0; i < n; i++ {
		events[2*i] = [4]int{cars[i][0], 1, cars[i][3] - cars[i][2] + 1, cars[i][4]}
		events[2*i+1] = [4]int{cars[i][1], -1, cars[i][3] - cars[i][2] + 1, cars[i][4]}
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

	for _, event := range events {
		if event[1] == -1 {
			occupied -= event[2]
			nowCars--
			minCarIndex--
		} else {
			occupied += event[2]
			nowCars++
			minCarsList[minCarIndex] = event[3]
			minCarIndex++
		}

		if occupied == maxCars && nowCars == minCarsCnt {
			minCarsList = minCarsList[0:minCarsCnt]
			return minCarsList
		}
	}

	return minCarsList[:0]
}
