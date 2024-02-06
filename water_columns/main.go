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

	fmt.Println("Enter columns count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter columns list")
	columns := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		columns[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Got max volume: ", getWaterVolume(n, columns))
}

func getWaterVolume(n int, columns []int) int {
	volume := 0
	maxIndex := 0
	max := 0
	for i, val := range columns {
		if val > max {
			max = val
			maxIndex = i
		}
	}

	max = 0
	for i := 0; i < maxIndex; i++ {
		if columns[i] > max {
			max = columns[i]
			continue
		}
		volume += max - columns[i]
	}

	max = 0
	for i := n - 1; i > maxIndex; i-- {
		if columns[i] > max {
			max = columns[i]
			continue
		}
		volume += max - columns[i]
	}

	return volume
}
