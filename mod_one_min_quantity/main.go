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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	unique := make(map[int]int, n)
	for _, v := range a {
		unique[v]++
	}

	minCnt := n
	for v, cnt := range unique {
		curCnt := n - cnt
		if curV, ok := unique[v+1]; ok {
			curCnt -= curV
		}
		if curCnt < minCnt {
			minCnt = curCnt
		}
	}

	fmt.Println(minCnt)

}
