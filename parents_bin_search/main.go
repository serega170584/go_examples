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

	fmt.Println("Enter numbers count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter parents count")
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	params := [2]int{k, n}
	fmt.Println("Added parents count", lbinsearch(0, n, check, params))
}

func lbinsearch(l int, r int, check func(m int, params [2]int) bool, params [2]int) int {
	for l < r {
		m := (l + r) / 2
		if check(m, params) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func check(m int, params [2]int) bool {
	k, n := params[0], params[1]
	if 3*(k+m) >= n+m {
		return true
	}
	return false
}
