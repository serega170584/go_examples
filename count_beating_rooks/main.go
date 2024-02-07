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

	fmt.Println("Enter count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter coords")
	coords := make([][2]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		coords[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		coords[i][1], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Got pairs", countBeatingRooks(n, coords))
}

func countBeatingRooks(cnt int, coords [][2]int) int {
	columns := make(map[int]int, cnt)
	rows := make(map[int]int, cnt)

	for _, val := range coords {
		if _, ok := columns[val[0]]; !ok {
			columns[val[0]] = 0
		}
		columns[val[0]]++

		if _, ok := rows[val[1]]; !ok {
			rows[val[1]] = 0
		}
		rows[val[1]]++
	}

	pairsCnt := 0

	for _, v := range columns {
		pairsCnt += v - 1
	}

	for _, v := range rows {
		pairsCnt += v - 1
	}

	return pairsCnt
}
