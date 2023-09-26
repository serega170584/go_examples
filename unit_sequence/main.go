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
	cntStr := scanner.Text()
	cnt, _ := strconv.Atoi(cntStr)

	arr := make([]string, cnt+1)

	for i := 0; i < cnt; i++ {
		scanner.Scan()
		arr[i] = scanner.Text()
	}

	singleCnt := 0
	startPointerInd := 0
	lastPointerInd := 0
	for startPointerInd < cnt+1 {
		for arr[lastPointerInd] == "1" {
			lastPointerInd++
		}
		if singleCnt < lastPointerInd-startPointerInd {
			singleCnt = lastPointerInd - startPointerInd
		}
		lastPointerInd++
		startPointerInd = lastPointerInd
	}

	fmt.Println(singleCnt)
}
