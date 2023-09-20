package main

import (
	"fmt"
	"log"
	"sort"
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

//5
//1 2
//3 4
//5 6
//7 15
//16 17
//5
//1 2
//3 4
//6 10
//11 12
//14 17
//1 2
//3 4
//6 6
//7 10
//11 12
//14 15
//16 17

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

	segmentPointFinishes := make(map[int]bool, 2*mainCnt+2*secondaryCnt)
	segmentPoints := make([]int, 2*mainCnt+2*secondaryCnt)
	segmentPointsCnt := make(map[int]int, 2*mainCnt+2*secondaryCnt)

	for i, val := range mainSegment {
		segmentPointFinishes[val[0]] = false
		segmentPointFinishes[val[1]] = true
		segmentPoints[i*2] = val[0]
		segmentPoints[i*2+1] = val[1]
		segmentPointsCnt[val[0]]++
		segmentPointsCnt[val[1]]++
	}

	for i, val := range secondary {
		segmentPoints[(i+mainCnt)*2] = 1000000001
		segmentPoints[(i+mainCnt)*2+1] = 1000000001

		if _, ok := segmentPointFinishes[val[0]]; !ok {
			segmentPointFinishes[val[0]] = false
			segmentPoints[(i+mainCnt)*2] = val[0]
		}

		if _, ok := segmentPointFinishes[val[1]]; !ok {
			segmentPoints[(i+mainCnt)*2+1] = val[1]
		}

		segmentPointFinishes[val[1]] = true

		segmentPointsCnt[val[0]]++
		segmentPointsCnt[val[1]]++
	}

	sort.Ints(segmentPoints)

	startCnt := 0
	for i, val := range segmentPoints {
		if val == 1000000000 {
			break
		}

		if segmentPointFinishes[val] {
			startCnt -= segmentPointsCnt[val]
		} else {
			startCnt += segmentPointsCnt[val]
		}

		if segmentPointFinishes[val] && startCnt == 0 && segmentPointsCnt[val] == 2 {
			fmt.Printf("%d %d\n", segmentPoints[i-1], val)
		}

		if segmentPointFinishes[val] && startCnt == 1 {
			fmt.Printf("%d %d\n", segmentPoints[i-1], val)
		}

		if segmentPointFinishes[val] && startCnt == -1 {
			startCnt = 1
			fmt.Printf("%d %d\n", val, val)
		}
	}
}
