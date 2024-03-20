package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func main() {
	scanner := makeScanner()

	scanner.Scan()
	a := []rune(scanner.Text())

	scanner.Scan()
	b := []rune(scanner.Text())

	isSame := isSameAnagram(a, b)
	if isSame {
		fmt.Println("YES")
		return
	}

	fmt.Println("NO")
}

func isSameAnagram(a []rune, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	ac := make(map[rune]int, len(a))
	bc := make(map[rune]int, len(b))

	al := make([]rune, 0, len(a))
	bl := make([]rune, 0, len(b))

	for _, v := range a {
		if _, ok := ac[v]; !ok {
			al = append(al, v)
		}
		ac[v]++
	}

	for _, v := range b {
		if _, ok := bc[v]; !ok {
			bl = append(bl, v)
		}
		bc[v]++
	}

	if len(al) != len(bl) {
		return false
	}

	l := len(al)

	slices.Sort(al)
	slices.Sort(bl)

	for i := 0; i < l; i++ {
		if al[i] != bl[i] {
			return false
		}

		if ac[al[i]] != bc[bl[i]] {
			return false
		}
	}

	return true
}
