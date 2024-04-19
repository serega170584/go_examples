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
	n, _ := strconv.Atoi(scanner.Text())

	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(getMinSyms(n, list))
}

func getMinSyms(n int, list []int) int {
	syms := make([]int, 4)
	syms[1] = 1
	syms[2] = 2
	syms[3] = 2

	q := 0
	for _, v := range list {
		q += v/4 + syms[v%4]
	}

	return q
}
