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

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	isRepeated := isNumberRepeated(n, k, list)
	if isRepeated {
		fmt.Println("YES")
		return
	}

	fmt.Println("NO")
}

func isNumberRepeated(n int, k int, list []int) bool {
	for i, v := range list {
		m := make(map[int]struct{}, k)
		m[v] = struct{}{}
		for j := i + 1; j <= i+k && j < n; j++ {
			if _, ok := m[list[j]]; ok {
				return true
			}
			m[list[j]] = struct{}{}
		}
	}
	return false
}
