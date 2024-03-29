package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	attempts := make([][2]int, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		attempts[i][0], _ = strconv.Atoi(scanner.Text())

		scanner.Scan()
		attempts[i][1], _ = strconv.Atoi(scanner.Text())
	}

	v := getStart(n, list, m, attempts)
	for i := 0; i < m; i++ {
		if v[i] == -1 {
			fmt.Println(v[i])
		} else {
			fmt.Println(v[i] + 1)
		}
	}
}

func getStart(n int, list []int, m int, attempts [][2]int) []int {
	sums := make([]int, n)
	for i, v := range list {
		if i == 0 {
			sums[i] = list[i]
		} else {
			sums[i] = sums[i-1] + v
		}
	}

	start := make([]int, 0, m)
	for _, v := range attempts {
		start = append(start, binarySearch(n, list, sums, v[0], v[1]))
	}
	return start
}

func binarySearch(n int, list []int, sums []int, cnt int, sum int) int {
	l := 0
	r := n - 1
	for l < r {
		m := (l + r) / 2
		if check(m, n, list, sums, cnt, sum) {
			r = m
		} else {
			l = m + 1
		}
	}
	if list[l]+sums[l+cnt-1]-sums[l] != sum {
		return -1
	}
	return l
}

func check(m int, n int, list []int, sums []int, cnt int, sum int) bool {
	if m+cnt-1 > n-1 {
		return true
	}
	if list[m]+sums[m+cnt-1]-sums[m] >= sum {
		return true
	}
	return false
}
