package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Params struct {
	x    int
	list []int
}

func NewParams(x int, list []int) *Params {
	return &Params{x: x, list: list}
}

func (params *Params) X() int {
	return params.x
}

func (params *Params) List() []int {
	return params.list
}

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

	fmt.Println("Enter search number")
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())

	params := NewParams(x, list)
	fmt.Println("Got searched index", lsearch(0, n-1, n, check, params))
}

func lsearch(l int, r int, n int, check func(m int, params *Params) bool, params *Params) int {
	for l < r {
		m := (l + r) / 2
		if check(m, params) {
			r = m
		} else {
			l = m + 1
		}
	}
	if params.List()[l] < params.X() {
		return n
	}
	return l
}

func check(m int, params *Params) bool {
	return params.List()[m] >= params.X()
}
