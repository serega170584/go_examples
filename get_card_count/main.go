package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 5 11 21 31 10 5 5 5 5 5
// 20 11 21 31 41 36 35 5
func cardCount(array []int, cnt int, stepCnt int) int {
	rightCards := make([]int, stepCnt)
	var maxCnt, rightInd int
	for i := cnt - stepCnt; i < cnt; i++ {
		val := array[i]
		maxCnt += val
		rightCards[rightInd] = val
		rightInd++
	}

	val := maxCnt
	for i := 0; i < stepCnt; i++ {
		val = val + array[i] - rightCards[i]
		if val > maxCnt {
			maxCnt = val
		}
	}

	return maxCnt
}

func main() {
	scanner := makeScanner()
	cnt := readInt(scanner)
	stepCnt := readInt(scanner)
	array := readArray(scanner)

	fmt.Print(cardCount(array, cnt, stepCnt))
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
