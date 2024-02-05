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

	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Got min even")
	fmt.Println(findMinEven(n, list))
}

func findMinEven(n int, list []int) int {
	isMet := false
	minVal := -1
	for _, v := range list {
		if v%2 == 1 {
			continue
		}
		if isMet {
			minVal = min(minVal, v)
		} else {
			minVal = v
			isMet = true
		}
	}
	return minVal
}
