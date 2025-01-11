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

	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	maxLenInd := 0
	maxVal := 0
	for i, v := range list {
		if v > maxVal {
			maxLenInd = i
			maxVal = v
		}
	}

	sum := 0
	for i, v := range list {
		if i != maxLenInd {
			sum += v
		}
	}

	rest := 0
	if sum < maxVal {
		rest = maxVal - sum
	} else {
		rest = sum + maxVal
	}

	fmt.Println(rest)
}
