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

	fmt.Println("Enter search number")
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Searched number index ", getNumberIndex(x, list))
}

func getNumberIndex(x int, list []int) int {
	for i, v := range list {
		if v == x {
			return i
		}
	}
	return -1
}
