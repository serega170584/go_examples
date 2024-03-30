package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	s := make([]string, 2)
	numerator, denominator := getKantorNum(n)
	s[0] = strconv.Itoa(numerator)
	s[1] = strconv.Itoa(denominator)

	fmt.Println(strings.Join(s, "/"))
}

func getKantorNum(n int) (int, int) {
	l := 1
	r := 2000000000
	for l < r {
		mid := (l + r) / 2
		if check(mid, n) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	part := (l-1)*l/2 + 1
	numerator := 1 + n - part
	denominator := l - n + part
	if l%2 == 0 {
		numerator, denominator = denominator, numerator
	}

	return numerator, denominator
}

func check(mid int, n int) bool {
	return mid*(mid+1)/2 >= n
}
