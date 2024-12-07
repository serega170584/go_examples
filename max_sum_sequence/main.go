package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	first := make([]int, m+n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		first[i], _ = strconv.Atoi(scanner.Text())
	}

	second := make([]int, m+n)
	for i := 0; i < m; i++ {
		scanner.Scan()
		second[i], _ = strconv.Atoi(scanner.Text())
	}

	indexes := make([]int, m+n)
	for i := 0; i < m+n; i++ {
		indexes[i] = i
	}

	i := 0
	j := 0

	sort.Slice(indexes, func(k int, l int) bool {
		if i == n && second[k] < second[l] {
			j++
			if second[k] < second[l] {
				return true
			} else {
				return false
			}
		}

		if j == m {
			i++
			if first[k] < first[l] {
				return true
			} else {
				return false
			}
		}

		return true
	})
}
