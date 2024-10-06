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
	cnt, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())

	list := make([]int, 0, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		list = append(list, v)
	}

	l := 0
	r := cnt - 1
	for l < r {
		m := (l + r) / 2
		if lCheck(list, m, s) {
			r = m
		} else {
			l = m + 1
		}
	}

	if list[r] != s {
		return
	}
	left := r

	l = 0
	r = cnt - 1
	for l < r {
		m := (l + r + 1) / 2
		if rCheck(list, m, s) {
			l = m
		} else {
			r = m - 1
		}
	}

	right := l

	fmt.Println(left, right)
}

func lCheck(list []int, m int, s int) bool {
	if s <= list[m] {
		return true
	}
	return false
}

func rCheck(list []int, m int, s int) bool {
	if s >= list[m] {
		return true
	}
	return false
}
