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

	fmt.Println("Enter list count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list")
	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	l, r := 0, n-1
	fmt.Println("Got searched index", lbinsearch(l, r, check, list))
}

func lbinsearch(l int, r int, check func(m int, list []int) bool, list []int) int {
	for l < r {
		m := (l + r) / 2
		if check(m, list) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func check(m int, list []int) bool {
	if 10 <= list[m] {
		return true
	}
	return false
}
