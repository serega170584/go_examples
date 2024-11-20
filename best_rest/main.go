package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	slices.Sort(list)

	interval := make([]int, n-1)
	for i := 1; i < n; i++ {
		interval[i-1] = list[i] - list[i-1]
	}

	sum := make([]int, n)
	sum[0] = 0
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + interval[i-1]
	}

	left := 0
	cnt := 1

	for i := 1; i < n; i++ {
		if sum[i]-sum[left] <= k {
			cnt++
			continue
		}
		left++
	}

	fmt.Println(cnt)
}
