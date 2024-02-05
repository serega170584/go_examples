package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter string counter")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter dictionary length")
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter string")
	scanner.Scan()
	s := scanner.Text()

	fmt.Println(maxCounter(n, k, s))
}

func maxCounter(n int, k int, s string) int {
	dict := make(map[rune]int, k)
	maxCount := 0

	for _, v := range s {
		if _, ok := dict[v]; !ok {
			dict[v] = 0
		}
		dict[v]++
		maxCount = max(maxCount, dict[v])
	}

	return maxCount
}
