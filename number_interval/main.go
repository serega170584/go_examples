package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	intervals := make([]int, 0)
	for i := 0; i < n; i++ {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		intervals = append(intervals, val)
	}

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	list := make([][2]int, k)
	for i := 0; i < k; i++ {
		scanner.Scan()
		list[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		list[i][1], _ = strconv.Atoi(scanner.Text())
	}

	lengths := getIntervalLengths(n, list, intervals)
	s := make([]string, n)
	for i, v := range lengths {
		s[i] = strconv.Itoa(v)
	}

	fmt.Println(strings.Join(s, " "))
}

func getIntervalLengths(n int, list [][2]int, intervals []int) []int {
	slices.Sort(intervals)

	lengths := make([]int, 0)
	for _, v := range list {
		l := binarySearch(v[0], n, intervals)
		r := rbinarySearch(v[1], n, intervals)
		length := 0
		if r >= l {
			length = r - l + 1
		}
		lengths = append(lengths, length)
	}

	return lengths
}

func binarySearch(x int, n int, list []int) int {
	l := 0
	r := n - 1
	m := 0
	for l < r {
		m = (l + r) / 2
		if check(m, x, list) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func rbinarySearch(x int, n int, list []int) int {
	l := 0
	r := n - 1
	for l < r {
		m := (l + r + 1) / 2
		if lcheck(m, x, list) {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}

func check(m int, x int, list []int) bool {
	return list[m] >= x
}

func lcheck(m int, x int, list []int) bool {
	return list[m] <= x
}
