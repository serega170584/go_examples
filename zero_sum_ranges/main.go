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

	fmt.Println("Enter list count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list")
	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	prefixSums := prefixZeros(n, list)
	dict := countPrefixSums(n, prefixSums)
	fmt.Println("Got intervals ", zeroIntervals(dict))
}

func prefixZeros(n int, list []int) []int {
	prefixSums := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		prefixSums[i] = prefixSums[i-1] + list[i-1]
	}
	return prefixSums
}

func countPrefixSums(n int, prefixSums []int) map[int]int {
	prefixSums = prefixSums[1:]
	dict := make(map[int]int, n)
	for _, val := range prefixSums {
		if _, ok := dict[val]; !ok {
			dict[val] = 0
		}
		dict[val]++
	}
	return dict
}

func zeroIntervals(dict map[int]int) int {
	cnt := 0
	for _, v := range dict {
		cnt += v * (v - 1) / 2
	}
	return cnt
}
