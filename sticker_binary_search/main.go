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

	fmt.Println("Enter first multiplier sticker number")
	scanner.Scan()
	m1, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter second multiplier sticker number")
	scanner.Scan()
	m2, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter width")
	scanner.Scan()
	w, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter height")
	scanner.Scan()
	h, _ := strconv.Atoi(scanner.Text())

	r := (w/m1 + 1) * (h/m2 + 1)
	n := m1 * m2
	params := [3]int{n, w, h}
	fmt.Println("Got sticker size", stickerBinarySearch(0, r, check, params))
}

func stickerBinarySearch(l int, r int, check func(m int, params [3]int) bool, params [3]int) int {
	for l < r {
		m := (l + r + 1) / 2
		if check(m, params) {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}

func check(m int, params [3]int) bool {
	n, w, h := params[0], params[1], params[2]
	wCnt := w / m
	hCnt := h / m
	return wCnt*hCnt >= n
}
