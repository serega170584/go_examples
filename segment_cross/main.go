package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// main segment and secondary segment
// 2 pointers: main and secondary
// look for nearest  end to start of segment: min(segment end, other one end)
// 1. segment end < other segment start = move first one pointer, stay pointer of other one
// 2. segment end > other segment end = move other one pointer, stay pointer of first one
// 3. cross = [other segment start, min segment end], move segment
// pointer > valid pointer = break cycle
// [1, 2] - [1, 3] - 3
// [1, 3] - [2, 3] - 3
// [1, 2] - [1, 2] - 3
// [1, 2] - [3, 4] - 1
// [1, 3] - [2, 4] - 3
// [1, 4] - [2, 3] - 2
// [1, 3] - [3, 4]

// [1, 2], [3, 5]  - [1, 4]
// [1, 2] - [1, 4], [5, 6]
// [1, 2], [3, 4] - [1, 4], [5, 6]
// [1, 2], [5, 8] - [1, 4], [5, 6]

// pointer move:
//

//decomposition:
// look for cross
// pointer move

// variables:
// main segment
// secondary segment
// leastStartSegment
// mostStartSegment
// cross array
// min segment end
func main() {
	var mainCnt, secondaryCnt int

	_, err := fmt.Scan(&mainCnt)
	if err != nil {
		log.Fatal(err)
	}

	scanner := makeScanner()

	mainSegment := make([][]int, mainCnt)

	for i := range mainSegment {
		scanner.Scan()
		mainStr := strings.Split(scanner.Text(), " ")
		mainSegment[i] = make([]int, 2)
		mainSegment[i][0], _ = strconv.Atoi(mainStr[0])
		mainSegment[i][1], _ = strconv.Atoi(mainStr[1])
	}

	_, err = fmt.Scan(&secondaryCnt)
	if err != nil {
		log.Fatal(err)
	}

	secondary := make([][]int, secondaryCnt)

	for i := range secondary {
		scanner.Scan()
		secondaryStr := strings.Split(scanner.Text(), " ")
		secondary[i] = make([]int, 2)
		secondary[i][0], _ = strconv.Atoi(secondaryStr[0])
		secondary[i][1], _ = strconv.Atoi(secondaryStr[1])
	}

	mainInd := 0
	secondaryInd := 0
	result := make([][]int, 0)
	for mainInd != mainCnt && secondaryInd != secondaryCnt {
		least := mainSegment[mainInd]
		most := secondary[secondaryInd]
		leastInd := &mainInd
		mostInd := &secondaryInd
		if least[0] > most[0] {
			least, most = most, least
			leastInd, mostInd = mostInd, leastInd
		}

		if least[1] < most[0] {
			*leastInd++
			continue
		}

		if least[1] > most[1] {
			*mostInd++
			el := make([]int, 2)
			el[0] = most[0]
			el[1] = most[1]
			result = append(result, el)
			continue
		}

		el := make([]int, 2)
		el[0] = most[0]
		el[1] = least[1]
		result = append(result, el)
		*leastInd++
	}

	for _, val := range result {
		fmt.Printf("%d %d\n", val[0], val[1])
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}
