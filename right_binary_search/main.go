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
	scanner.Scan()
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Got index", rbinsearch(0, n-1, check, list))
}

func rbinsearch(l int, r int, check func(m int, list []int) bool, list []int) int {
	for l < r {
		m := (l + r + 1) / 2
		if check(m, list) {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}

func check(m int, list []int) bool {
	if list[m] <= 10 {
		return true
	}
	return false
}
