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

func (p *Params) List() []int {
	return p.list
}

func (p *Params) X() int {
	return p.x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter count")
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
	fmt.Println("Searched count", countx(n, params))
}

func lbinarysearch(l int, r int, check func(m int, params *Params) bool, params *Params) int {
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

func findFirst(n int, params *Params, check func(m int, params *Params) bool) int {
	searchedIndex := lbinarysearch(0, n-1, check, params)
	if !check(searchedIndex, params) {
		return n
	}
	return searchedIndex
}

func countx(n int, params *Params) int {
	eqNumberIndex := findFirst(n, params, checkEq)
	moreNumberIndex := findFirst(n, params, checkMore)
	return moreNumberIndex - eqNumberIndex
}

func checkEq(m int, params *Params) bool {
	return params.List()[m] >= params.X()
}

func checkMore(m int, params *Params) bool {
	return params.List()[m] > params.X()
}
