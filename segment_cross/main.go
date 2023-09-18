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
func main() {
	var mainCnt, secondaryCnt int

	_, err := fmt.Scan(&mainCnt)
	if err != nil {
		log.Fatal(err)
	}

	segments := make([][]int, 2)

	segments[0] = make([]int, 2*mainCnt)

	for i := 0; i < mainCnt; i++ {
		var start, finish string
		_, _ = fmt.Scan(&start, &finish)
		segments[0][2*i], _ = strconv.Atoi(start)
		segments[0][2*i+1], _ = strconv.Atoi(finish)
	}

	_, err = fmt.Scan(&secondaryCnt)
	if err != nil {
		log.Fatal(err)
	}

	segments[1] = make([]int, 2*secondaryCnt)

	for i := 0; i < secondaryCnt; i++ {
		var start, finish string
		_, _ = fmt.Scan(&start, &finish)
		segments[1][2*i], _ = strconv.Atoi(start)
		segments[1][2*i+1], _ = strconv.Atoi(finish)
	}

	segmentPointer := make([]int, 2)

	segmentInd := 0

	prevFinish := segments[segmentInd][0]
	finish := segments[segmentInd][1]

	segmentInd = 1

	for (segmentInd == 0 && mainCnt != segmentPointer[0]) || (segmentInd == 1 && segmentPointer[1] != secondaryCnt) {
		pointerInd := segmentPointer[segmentInd]

		isContinue := false

		for segments[segmentInd][2*pointerInd+1] < prevFinish {
			segmentPointer[segmentInd]++

			pointerInd = segmentPointer[segmentInd]

			if (segmentInd == 0 && mainCnt == segmentPointer[0]) || (segmentInd == 1 && segmentPointer[1] == secondaryCnt) {
				isContinue = true
				break
			}
		}

		if isContinue {
			continue
		}

		for segments[segmentInd][2*pointerInd] > finish {
			isContinue = true
			prevFinish = segments[segmentInd][2*pointerInd]
			finish = segments[segmentInd][2*pointerInd+1]
			segmentInd = (segmentInd + 1) % 2
			segmentPointer[segmentInd]++

			pointerInd = segmentPointer[segmentInd]

			if (segmentInd == 0 && mainCnt == segmentPointer[0]) || (segmentInd == 1 && segmentPointer[1] == secondaryCnt) {
				break
			}
		}

		if isContinue {
			continue
		}

		for prevFinish < segments[segmentInd][2*pointerInd+1] && segments[segmentInd][2*pointerInd+1] < finish {
			maxStart := prevFinish
			if segments[segmentInd][2*pointerInd] > maxStart {
				maxStart = segments[segmentInd][2*pointerInd]
			}

			isContinue = true

			prevFinish = segments[segmentInd][2*pointerInd+1]

			fmt.Printf("%d %d\n", maxStart, segments[segmentInd][2*pointerInd+1])

			segmentPointer[segmentInd]++

			pointerInd = segmentPointer[segmentInd]

			if (segmentInd == 0 && mainCnt == segmentPointer[0]) || (segmentInd == 1 && segmentPointer[1] == secondaryCnt) {
				break
			}
		}

		if isContinue {
			continue
		}

		for segments[segmentInd][2*pointerInd] < finish && finish < segments[segmentInd][2*pointerInd+1] {
			maxStart := prevFinish
			if segments[segmentInd][2*pointerInd] > maxStart {
				maxStart = segments[segmentInd][2*pointerInd]
			}
			fmt.Printf("%d %d\n", maxStart, finish)

			segmentPointer[segmentInd]++
			prevFinish = finish
			finish = segments[segmentInd][2*pointerInd+1]
			segmentInd = (segmentInd + 1) % 2

			pointerInd = segmentPointer[segmentInd]

			if (segmentInd == 0 && mainCnt == segmentPointer[0]) || (segmentInd == 1 && segmentPointer[1] == secondaryCnt) {
				break
			}
		}
	}
}
