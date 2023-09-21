package main

import (
	"bufio"
	"os"
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

//2
//1 4
//6 9
//3
//2 3
//6 7
//8 10
//2 3
//6 7
//8 9

func main() {
	//var mainCnt, secondaryCnt int

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Scan()
	mainCnt, _ := strconv.Atoi(s.Text())

	mainSegment := make([][]int, mainCnt)

	for i := 0; i < mainCnt; i++ {
		mainSegment[i] = make([]int, 2)
		s.Scan()
		mainSegment[i][0], _ = strconv.Atoi(s.Text())
		s.Scan()
		mainSegment[i][1], _ = strconv.Atoi(s.Text())
	}

	s.Scan()
	secondaryCnt, _ := strconv.Atoi(s.Text())

	for i := 0; i < secondaryCnt; i++ {
		mainSegment[i] = make([]int, 2)
		s.Scan()
		mainSegment[i][0], _ = strconv.Atoi(s.Text())
		s.Scan()
		mainSegment[i][1], _ = strconv.Atoi(s.Text())
	}

	//_, err := fmt.Scan(&mainCnt)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//mainSegment := make([][]int, mainCnt)
	//
	//for i := range mainSegment {
	//	var left, right string
	//	_, _ = fmt.Scan(&left, &right)
	//	mainSegment[i] = make([]int, 2)
	//	mainSegment[i][0], _ = strconv.Atoi(left)
	//	mainSegment[i][1], _ = strconv.Atoi(right)
	//}
	//
	//_, err = fmt.Scan(&secondaryCnt)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//secondary := make([][]int, secondaryCnt)
	//
	//for i := range secondary {
	//	var left, right string
	//	_, _ = fmt.Scan(&left, &right)
	//	secondary[i] = make([]int, 2)
	//	secondary[i][0], _ = strconv.Atoi(left)
	//	secondary[i][1], _ = strconv.Atoi(right)
	//}
	//
	//if mainCnt == 0 || secondaryCnt == 0 {
	//	return
	//}
	//
	//movePointer := 0
	//persistentPointer := 0
	//moveSegment := mainSegment
	//persistentSegment := secondary
	//moveCnt := mainCnt
	//persistentCnt := secondaryCnt
	//
	//for movePointer != moveCnt {
	//
	//	if persistentSegment[persistentPointer][1] < moveSegment[movePointer][1] {
	//		moveSegment, persistentSegment = persistentSegment, moveSegment
	//		movePointer, persistentPointer = persistentPointer, movePointer
	//		moveCnt, persistentCnt = persistentCnt, moveCnt
	//	}
	//
	//	if moveSegment[movePointer][0] >= persistentSegment[persistentPointer][0] {
	//		println(moveSegment[movePointer][0], moveSegment[movePointer][1])
	//		movePointer++
	//		continue
	//	}
	//
	//	if moveSegment[movePointer][1] >= persistentSegment[persistentPointer][0] {
	//		println(persistentSegment[persistentPointer][0], moveSegment[movePointer][1])
	//	}
	//
	//	movePointer++
	//}
}
