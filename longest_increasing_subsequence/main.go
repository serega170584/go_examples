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

	a := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(lis(a))
}

func lis(a []int) int {
	cnt := len(a)
	incLen := make([]int, cnt)
	incLen[0] = 1
	longLen := 1
	for i := 1; i < cnt; i++ {
		incLen[i] = 1
		for j := 0; j < i; j++ {
			if a[i] > a[j] && incLen[j]+1 > incLen[i] {
				incLen[i] = incLen[j] + 1
			}
		}
	}

	for i := 1; i < cnt; i++ {
		if incLen[i] > longLen {
			longLen = incLen[i]
		}
	}

	return longLen
}
