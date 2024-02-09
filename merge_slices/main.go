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

	fmt.Println("Enter first list count")
	scanner.Scan()
	fCnt, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter first list")
	fList := make([]int, fCnt)
	for i := 0; i < fCnt; i++ {
		scanner.Scan()
		fList[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Enter second list count")
	scanner.Scan()
	sCnt, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter second list")
	sList := make([]int, sCnt)
	for i := 0; i < sCnt; i++ {
		scanner.Scan()
		sList[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Got merged list", mergeList(fCnt, sCnt, fList, sList))
}

func mergeList(fCnt int, sCnt int, first []int, second []int) []int {
	first = append(first, make([]int, sCnt)...)
	fIndex := fCnt - 1
	sIndex := sCnt - 1
	for i := fCnt + sCnt - 1; i > -1; i-- {
		if fIndex == -1 || first[fIndex] < second[sIndex] {
			first[i] = second[sIndex]
			sIndex--
		} else {
			first[i] = first[fIndex]
			fIndex--
		}
	}

	return first
}
