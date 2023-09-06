package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//1 4 5 4
//2 3 6 5
//3 4 7 8
//4 5 6 9

func main() {
	scanner := makeScanner()
	rowsCnt, colsCnt := readSize(scanner)
	input := readArr(scanner, rowsCnt, colsCnt)

	output := makeArr(rowsCnt, colsCnt)
	maxPath := 0

	for i := 0; i < rowsCnt; i++ {
		for j := 0; j < colsCnt; j++ {
			output = makeWeights(input, output, i, j, rowsCnt, colsCnt)
			if output[i][j] > maxPath {
				maxPath = output[i][j]
			}
		}
	}

	fmt.Println(maxPath)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readSize(scanner *bufio.Scanner) (int, int) {
	scanner.Scan()
	size := make([]int, 2)
	list := strings.Split(scanner.Text(), " ")
	i := 0
	for _, val := range list {
		size[i], _ = strconv.Atoi(val)
		i++
	}
	return size[0], size[1]
}

func readArr(scanner *bufio.Scanner, rowsCnt int, colsCnt int) [][]int {
	arr := make([][]int, rowsCnt)
	for i := 0; i < rowsCnt; i++ {
		scanner.Scan()
		arr[i] = make([]int, colsCnt)
		str := strings.Split(scanner.Text(), " ")
		for j := 0; j < colsCnt; j++ {
			arr[i][j], _ = strconv.Atoi(str[j])
		}
	}
	return arr
}

func makeArr(rowsCnt, colsCnt int) [][]int {
	arr := make([][]int, rowsCnt)
	for i := 0; i < rowsCnt; i++ {
		arr[i] = make([]int, colsCnt)
		for j := 0; j < colsCnt; j++ {
			arr[i][j] = -1
		}
	}
	return arr
}

func makeWeights(input [][]int, output [][]int, i int, j int, rowsCnt int, colsCnt int) [][]int {
	if output[i][j] != -1 {
		return output
	}

	weight := -1
	if i != 0 && input[i][j] > input[i-1][j] {
		output = makeWeights(input, output, i-1, j, rowsCnt, colsCnt)
		weight = output[i-1][j] + 1
		output[i][j] = weight
	}

	if i != rowsCnt-1 && input[i][j] > input[i+1][j] {
		output = makeWeights(input, output, i+1, j, rowsCnt, colsCnt)
		weight = output[i+1][j] + 1
	}

	if weight > output[i][j] {
		output[i][j] = weight
	}

	if j != 0 && input[i][j] > input[i][j-1] {
		output = makeWeights(input, output, i, j-1, rowsCnt, colsCnt)
		weight = output[i][j-1] + 1
	}

	if weight > output[i][j] {
		output[i][j] = weight
	}

	if j != colsCnt-1 && input[i][j] > input[i][j+1] {
		output = makeWeights(input, output, i, j+1, rowsCnt, colsCnt)
		weight = output[i][j+1] + 1
	}

	if weight > output[i][j] {
		output[i][j] = weight
	}

	if weight == -1 {
		output[i][j] = 1
	}

	return output
}
