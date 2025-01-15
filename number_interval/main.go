package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	scanner.Scan()
	requestsCnt, _ := strconv.Atoi(scanner.Text())

	requests := make([][2]int, requestsCnt)
	for i := 0; i < requestsCnt; i++ {
		scanner.Scan()
		requests[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		requests[i][1], _ = strconv.Atoi(scanner.Text())
	}

	slices.Sort(list)

	searchedCnt := make([]int, requestsCnt)
	for i, request := range requests {
		from := request[0]
		to := request[1]
		for v := from; v <= to; v++ {
			leftIndex := leftSearch(list, v)
			if leftIndex == -1 {
				continue
			}
			rightIndex := rightSearch(list, v)
			searchedCnt[i] += rightIndex - leftIndex + 1
		}
	}

	printSearched := make([]any, requestsCnt)
	for i, v := range searchedCnt {
		printSearched[i] = v
	}

	fmt.Println(printSearched...)
}

func leftSearch(list []int, v int) int {
	l := 0
	r := len(list) - 1
	m := 0
	for l < r {
		m = (l + r) / 2
		if list[m] >= v {
			r = m
		} else {
			l = m + 1
		}
	}
	if list[r] == v {
		return r
	}

	return -1
}

func rightSearch(list []int, v int) int {
	l := 0
	r := len(list) - 1
	m := 0
	for l < r {
		m = (l + r + 1) / 2
		if list[m] <= v {
			l = m
		} else {
			r = m - 1
		}
	}
	if list[l] == v {
		return l
	}
	return -1
}
