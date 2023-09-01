package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func status200Cnt(array []int, cnt int) int {
	var pairsCnt int

	modCounts := make([]int, 200)

	for i := 0; i < cnt; i++ {
		mod := array[i] % 200
		modCounts[mod]++
	}

	for _, cnt := range modCounts {
		pairsCnt += cnt * (cnt - 1) / 2
	}

	return pairsCnt
}

func main() {
	scanner := makeScanner()
	cnt := readInt(scanner)
	array := readArray(scanner)

	fmt.Print(status200Cnt(array, cnt))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for ind, val := range listString {
		arr[ind], _ = strconv.Atoi(val)
	}
	return arr
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
