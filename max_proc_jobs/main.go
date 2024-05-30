package main

import "fmt"

func main() {
	fmt.Println(getMaxProcJobs(4, 13, 4))
}

// 4 3 2 1
// 5 4 3 2
// 0 1 2 ... procCnt - 1
// proc - 1
// 1 2 3| 4 5 6 7
// 3 - 1  7 - 3
// 7 - 3 - 3 + 1 = 2
// 3 4 5 4 3 2 1
//
// 1 2 3 4 5| 6 7
// 5 - 1 7 - 5
// 7 - 5 - 5 + 1 = -2
// 1 2 3 4 5 4 3
// 2 3 4
// (4 + 2) * (4 - 2 + 1) / 2
func getMaxProcJobs(procCnt int, jobCnt int, proc int) int {
	leftDiff := proc - 1
	rightDiff := procCnt - proc
	diff := 0
	maxProcCnt := 1
	if leftDiff > rightDiff {
		diff = leftDiff - rightDiff
		maxProcCnt += leftDiff
	} else {
		diff = rightDiff - leftDiff
		maxProcCnt += rightDiff
	}
	firstCorner := 1
	secondCorner := 1 + diff
	availableJobsCnt := (firstCorner+maxProcCnt)*(maxProcCnt-firstCorner+1)/2 + (secondCorner+maxProcCnt)*(maxProcCnt-secondCorner+1)/2 - maxProcCnt
	if jobCnt > availableJobsCnt {
		maxProcCnt = (jobCnt-availableJobsCnt)/procCnt + maxProcCnt
	} else if jobCnt < availableJobsCnt {
		maxProcCnt -= (availableJobsCnt-jobCnt)/procCnt + 1
	}
	return maxProcCnt
}
