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

	fmt.Println("Enter from index")
	scanner.Scan()
	l, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter to index")
	scanner.Scan()
	r, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list")
	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	prefixZeros := PrefixZeros(n, list)
	fmt.Println("Got prefix zeros ", countZeros(prefixZeros, l, r))
}

func PrefixZeros(n int, list []int) []int {
	prefixZeros := make([]int, n+1)
	for i := 1; i < n; i++ {
		prefixZeros[i] = prefixZeros[i-1]
		if list[i-1] == 0 {
			prefixZeros[i]++
		}
	}
	return prefixZeros
}

func countZeros(prefixZeros []int, l int, r int) int {
	return prefixZeros[r] - prefixZeros[l]
}
