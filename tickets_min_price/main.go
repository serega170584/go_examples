package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter A")
	a := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Enter B")
	b := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Enter C")
	c := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		c[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Got min price", minPrice(n, a, b, c))
}

func minPrice(n int, a []int, b []int, c []int) int {
	minPrices := make([]int, n)
	for i := 0; i < n; i++ {
		minPrices[i] = math.MaxInt
	}

	minPrices[0] = a[0]

	for i := 0; i < n; i++ {
		prevC := i - 3
		if prevC >= 0 {
			minPrices[i] = min(minPrices[prevC]+c[i-2], minPrices[i])
		}

		prevB := i - 2
		if prevB >= 0 {
			minPrices[i] = min(minPrices[prevB]+b[i-2], minPrices[i])
		}

		prevA := i - 1
		if prevA >= 0 {
			minPrices[i] = min(minPrices[prevA]+a[i-1], minPrices[i])
		}
	}

	return minPrices[n-1]
}
