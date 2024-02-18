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

	fmt.Println("Enter steps count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter steps max size")
	scanner.Scan()
	stepsMaxSize, _ := strconv.Atoi(scanner.Text())

	wc := waysCnt(n, stepsMaxSize)
	fmt.Println("Got ways count", wc[n])
}

func waysCnt(n int, stepsMaxSize int) []int {
	ways := make([]int, n+1)
	ways[0] = 1
	ways[1] = 1
	for i := 2; i < n+1; i++ {
		for j := 0; j < stepsMaxSize; j++ {
			prev := i - j - 1
			if prev >= 0 {
				ways[i] += ways[prev]
			}
		}
	}
	return ways
}
