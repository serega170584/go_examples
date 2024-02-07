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

	fmt.Println("Sorted list ", countSort(n, list))
}

func countSort(cnt int, list []int) []int {
	minVal := 20000
	for _, v := range list {
		minVal = min(minVal, v)
	}
	countList := make([]int, cnt)
	for _, val := range list {
		key := val - minVal
		countList[key]++
	}
	valsList := make([]int, cnt)
	ind := 0
	for i, val := range countList {
		for j := 0; j < val; j++ {
			valsList[ind] = i + minVal
			ind++
		}
	}

	return valsList
}
