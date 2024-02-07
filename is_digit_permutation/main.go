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

	fmt.Println("Enter a")
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter b")
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Is digit permutation", isDigitPermutation(n, a, b))
}

func isDigitPermutation(cnt int, a int, b int) bool {
	aCounts := [10]int{}
	for i := 0; i < cnt; i++ {
		key := a % 10
		a = a / 10
		aCounts[key]++
	}
	bCounts := [10]int{}
	for i := 0; i < cnt; i++ {
		key := b % 10
		b = b / 10
		bCounts[key]++
	}
	return aCounts == bCounts
}
