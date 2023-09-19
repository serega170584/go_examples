package main

import (
	"fmt"
	"log"
	"strconv"
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

// [1 2]
// [3 4]
// 1 - 3 < 0, 1 - 4 < 0, 2 - 3 <= 0, 2 - 4 < 0 0000 0
// 3 - 1 > 0, 3 - 2 > 0, 4 - 1 > 0, 4 - 2 > 0 1111 15

// [1 3]
// [2 4]
// 1 - 2 < 0, 1 - 4 < 0, 3 - 2 > 0, 3 - 4 < 0 0010 2
// 2 - 1 > 0, 2 - 3 < 0, 4 - 1 > 0, 4 - 3 > 0 1011 11

// [2 3]
// [1 4]
// 2 - 1 > 0, 2 - 4 < 0, 3 - 1 > 0, 3 - 4 < 0 1010 10
// 1 - 2 < 0, 1 - 3 < 0, 4 - 2 > 0, 4 - 3 > 0 0011 3

// 1 3
// 2 4
// 1 2 3 4
// 1 - 2
// 3 - 2
// 3 - 4
func main() {
	var mainCnt, secondaryCnt int

	_, err := fmt.Scan(&mainCnt)
	if err != nil {
		log.Fatal(err)
	}

	mainSegment := make([][]int, mainCnt)

	for i := range mainSegment {
		var left, right string
		_, _ = fmt.Scan(&left, &right)
		mainSegment[i] = make([]int, 2)
		mainSegment[i][0], _ = strconv.Atoi(left)
		mainSegment[i][1], _ = strconv.Atoi(right)
	}

	_, err = fmt.Scan(&secondaryCnt)
	if err != nil {
		log.Fatal(err)
	}

	secondary := make([][]int, secondaryCnt)

	for i := range secondary {
		var left, right string
		_, _ = fmt.Scan(&left, &right)
		secondary[i] = make([]int, 2)
		secondary[i][0], _ = strconv.Atoi(left)
		secondary[i][1], _ = strconv.Atoi(right)
	}

	mainInd := 0
	secondaryInd := 0
	intBoolMapping := make(map[bool]int, 2)
	intBoolMapping[true] = 1
	intBoolMapping[false] = 0

	for mainInd != mainCnt && secondaryInd != secondaryCnt {
		diff := mainSegment[mainInd][1] - secondary[secondaryInd][0]
		finishDiff := mainSegment[mainInd][1] - secondary[secondaryInd][1]
		startDiff := mainSegment[mainInd][0] - secondary[secondaryInd][0]

		mainInd = mainInd + intBoolMapping[diff < 0]
		if diff < 0 {
			continue
		}

		if secondary[secondaryInd][1] < mainSegment[mainInd][0] {
			secondaryInd++
			continue
		}

		diff = diff - startDiff*intBoolMapping[startDiff > 0]
		start := intBoolMapping[startDiff <= 0]*secondary[secondaryInd][0] + intBoolMapping[startDiff > 0]*mainSegment[mainInd][0]

		mainInd = mainInd + intBoolMapping[finishDiff == 0]
		secondaryInd = secondaryInd + intBoolMapping[finishDiff == 0]

		diff = diff - intBoolMapping[finishDiff > 0]*finishDiff
		secondaryInd = secondaryInd + intBoolMapping[finishDiff > 0]

		mainInd = mainInd + intBoolMapping[finishDiff < 0]

		fmt.Printf("%d %d\n", start, start+diff)
	}
}
