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
	cnt1, _ := strconv.Atoi(scanner.Text())

	a := make([]int, cnt1)
	for i := 0; i < cnt1; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	scanner.Scan()
	cnt2, _ := strconv.Atoi(scanner.Text())

	b := make([]int, cnt2)
	for i := 0; i < cnt2; i++ {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(diff(a, b))
}

func diff(a []int, b []int) []int {
	al := len(a)
	c := make([]int, 0, al)
	bi := 0
	bl := len(b)
	for _, v := range a {
		for bi != bl && b[bi] < v {
			bi++
		}

		if bi == bl || v != b[bi] {
			c = append(c, v)
		}
	}

	return c
}
