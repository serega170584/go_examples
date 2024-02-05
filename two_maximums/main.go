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

	max2, max1 := getMaximums(n, list)
	fmt.Println("Get max 2: ", max2, ", get max: ", max1)
}

func getMaximums(n int, list []int) (int, int) {
	max2, max1 := list[0], list[1]
	if max2 > max1 {
		max2, max1 = max1, max2
	}

	for i := 2; i < n; i++ {
		if list[i] > max1 {
			max2 = max1
			max1 = list[i]
		} else if list[i] > max2 {
			max2 = list[i]
		}
	}
	return max2, max1
}
