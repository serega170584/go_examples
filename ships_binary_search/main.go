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

	fmt.Println(binarySearch(n))
}

func getShipsCnt(k int) int {
	if k == 1 {
		return 1
	}

	cnt := 1
	for i := 2; i <= k; i++ {
		cnt = i*(i+1)/2 + i + cnt
	}

	return cnt
}

func binarySearch(n int) int {
	l := 1
	r := 2000000
	for l < r {
		m := (l + r) / 2
		if check(m, n) {
			r = m
		} else {
			l = m + 1
		}
	}
	if getShipsCnt(l) > n {
		l--
	}
	return l
}

func check(m int, n int) bool {
	return getShipsCnt(m) >= n
}
