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
	firstCnt, _ := strconv.Atoi(scanner.Text())
	firstList := make([]int, firstCnt)
	for i := 0; i < firstCnt; i++ {
		scanner.Scan()
		firstList[i], _ = strconv.Atoi(scanner.Text())
	}

	scanner.Scan()
	secondCnt, _ := strconv.Atoi(scanner.Text())
	secondList := make([]int, secondCnt)
	for i := 0; i < secondCnt; i++ {
		scanner.Scan()
		secondList[i], _ = strconv.Atoi(scanner.Text())
	}

	scanner.Scan()
	thirdCnt, _ := strconv.Atoi(scanner.Text())
	thirdList := make([]int, thirdCnt)
	for i := 0; i < thirdCnt; i++ {
		scanner.Scan()
		thirdList[i], _ = strconv.Atoi(scanner.Text())
	}

	firstUnique := make(map[int]struct{}, len(firstList))
	secondUnique := make(map[int]struct{}, len(secondList))
	thirdUnique := make(map[int]struct{}, len(thirdList))

	for i := 0; i < firstCnt; i++ {
		firstUnique[firstList[i]] = struct{}{}
	}

	for i := 0; i < secondCnt; i++ {
		secondUnique[secondList[i]] = struct{}{}
	}

	for i := 0; i < thirdCnt; i++ {
		thirdUnique[thirdList[i]] = struct{}{}
	}

	numCnt := make(map[int]int, firstCnt+secondCnt+thirdCnt)

	for v := range firstUnique {
		numCnt[v]++
	}

	for v := range secondUnique {
		numCnt[v]++
	}

	for v := range thirdUnique {
		numCnt[v]++
	}

	res := make([]any, 0, len(numCnt))
	for i, v := range numCnt {
		if v >= 2 {
			res = append(res, i)
		}
	}

	fmt.Println(res...)
}
