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

	spaceCounts := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		spaceCounts[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(getMinSpacesCnt(spaceCounts))
}

func getMinSpacesCnt(spaceCounts []int) int {
	baseSpaceCounts := make([]int, 4)
	baseSpaceCounts[1] = 1
	baseSpaceCounts[2] = 2
	baseSpaceCounts[3] = 2

	minSpacesCnt := 0
	for _, v := range spaceCounts {
		minSpacesCnt += v/4 + baseSpaceCounts[v%4]
	}

	return minSpacesCnt
}
