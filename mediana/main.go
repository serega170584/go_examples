package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	slices.Sort(list)

	mediana := make([]any, 0, n)

	left := -1
	right := -1
	if n%2 == 0 {
		left = n/2 - 1
		right = n / 2
	} else {
		left = n/2 - 1
		right = n/2 + 1
		mediana = append(mediana, list[n/2])
	}

	for left != -1 && right != n {
		mediana = append(mediana, list[left])
		left--
		mediana = append(mediana, list[right])
		right++
	}

	if left != -1 {
		mediana = append(mediana, list[left])
	}

	if right != n {
		mediana = append(mediana, list[right])
	}

	fmt.Println(mediana...)
}
