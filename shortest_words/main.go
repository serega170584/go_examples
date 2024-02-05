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
	list := make([][]int, n)
	for i := 0; i < n; i++ {
		fmt.Println("Enter word length")
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		list[i] = make([]int, m)
		for j := 0; j < m; j++ {
			scanner.Scan()
			list[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}

	fmt.Println("Shortest words: ", getShortestWords(n, list))
}

func getShortestWords(n int, list [][]int) [][]int {
	isMet := false
	minValue := 0
	shortestWords := make([][]int, n)
	for _, v := range list {
		if isMet {
			minValue = min(minValue, len(v))
		} else {
			minValue = len(v)
		}
	}

	shortestWordIndex := 0
	for _, v := range list {
		if minValue == len(v) {
			shortestWords[shortestWordIndex] = v
			shortestWordIndex++
		}
	}

	return shortestWords[0:shortestWordIndex]
}
