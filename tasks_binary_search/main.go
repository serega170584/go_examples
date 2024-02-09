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

	fmt.Println("Enter total tasks")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter solved tasks")
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Solved tasks days count", ltasksbinarysearch(0, n, check, [2]int{k, n}))
}

func ltasksbinarysearch(l int, r int, check func(m int, params [2]int) bool, params [2]int) int {
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
	return (2*k+m-1)*m/2 > n
}
