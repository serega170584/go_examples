package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//5 10 15 2 10 15 5 5 5 20 20 1 20 1 1

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list")
	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Got max subsequence length", maxGrowingSubsequence(n, list))
}

func maxGrowingSubsequence(n int, list []int) int {
	maxVal := 1
	maxList := make([]int, n)

	for i := range maxList {
		maxList[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if list[i] > list[j] {
				val := maxList[j] + 1
				if val > maxList[i] {
					maxList[i] = val
				}
			}
		}
		maxVal = max(maxList[i], maxVal)
	}

	fmt.Println(maxList)

	return maxVal
}
