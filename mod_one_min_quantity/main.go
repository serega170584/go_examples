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

	fmt.Println(getModOneMinCnt(list, n))
}

func getModOneMinCnt(list []int, n int) int {
	m := make(map[int]int, n)
	for _, v := range list {
		m[v]++
	}

	minExcludedCnt := n
	for v := range m {
		curCnt := n - m[v]

		if _, ok := m[v-1]; ok {
			minExcludedCnt = min(minExcludedCnt, curCnt-m[v-1])
		}

		if _, ok := m[v+1]; ok {
			minExcludedCnt = min(minExcludedCnt, curCnt-m[v+1])
		}
	}

	return minExcludedCnt
}
