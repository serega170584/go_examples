package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 0 1 2 3 4 5 6 5 4 3
// 0 1 2 3 4
// 4 * 5 / 2 = 10
// 1 2 3 4 5
// (n + 1) * n / 2 = 6 * 2 = 12
// (n + 1) * (n - 1) / 2 + (n + 1) / 2= (n+1) * ((n-1)/2+1/2) = (n+1)*(n/2)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	ind, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	limit, _ := strconv.Atoi(scanner.Text())

	h := n - 1 - ind
	l := ind
	if l >= h {
		l, h = h, l
	}

	s := 0
	if h != 0 {
		s += h * (h + 1) / 2
	}

	if l != 0 {
		s += (h + h - l) * (h - (h - l) + 1) / 2
		s -= h
	}

	if s == 0 {
		fmt.Println(limit)
		return
	}

	if s > limit {
		rest := (s-limit)/n + 1
		fmt.Println(h - rest)
	} else {
		rest := (limit - s) / n
		fmt.Println(rest + h)
	}
}
