package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type PathNode struct {
	prevI int
	prevJ int
}

type CurNode struct {
	i int
	j int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)

	fmt.Println("Enter first word count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter first word")
	firstWord := make([]rune, n)
	scanner.Scan()
	for i := 0; i < n; i++ {
		scanner.Scan()
		firstWord[i] = []rune(scanner.Text())[0]
	}

	scanner.Scan()

	fmt.Println("Enter second word count")
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter second word")
	secondWord := make([]rune, m)
	scanner.Scan()
	for i := 0; i < m; i++ {
		scanner.Scan()
		secondWord[i] = []rune(scanner.Text())[0]
	}

	maxLength, path := maxSharedSubsequence(n, m, firstWord, secondWord)

	fmt.Println("Got subsequence max length", maxLength)
	fmt.Println("Got max subsequence path")
	for _, v := range path {
		fmt.Println(*v)
	}
}

func maxSharedSubsequence(rowSymCnt int, colSymCnt int, rowString []rune, colString []rune) (int, []*CurNode) {
	path := make([][]*PathNode, rowSymCnt)
	pathListCnt := rowSymCnt + colSymCnt
	pathList := make([]*CurNode, pathListCnt)
	dp := make([][]int, rowSymCnt)
	for i := 0; i < rowSymCnt; i++ {
		dp[i] = make([]int, colSymCnt)
		path[i] = make([]*PathNode, colSymCnt)
		for j := 0; j < colSymCnt; j++ {
			prevI := i - 1
			prevJ := j - 1
			if rowString[i] == colString[j] {
				prevVal := 0
				if prevI >= 0 && prevJ >= 0 {
					prevVal = dp[prevI][prevJ]
				}
				dp[i][j] = prevVal + 1
				pathNode := &PathNode{prevI: prevI, prevJ: prevJ}
				path[i][j] = pathNode
			} else {
				pathNode := &PathNode{prevI: -1, prevJ: -1}

				prevIVal := 0
				if prevI >= 0 {
					prevIVal = dp[prevI][j]
				}

				prevJVal := 0
				if prevJ >= 0 {
					prevJVal = dp[i][prevJ]
				}

				if prevI >= 0 && prevIVal == max(prevIVal, prevJVal) {
					dp[i][j] = prevIVal
					*pathNode = PathNode{prevI: prevI, prevJ: j}
					path[i][j] = pathNode
				} else {
					dp[i][j] = prevJVal
					*pathNode = PathNode{prevI: i, prevJ: prevJ}
					path[i][j] = pathNode
				}
			}
		}
	}

	ind := pathListCnt - 1
	curI := rowSymCnt - 1
	curJ := colSymCnt - 1
	for curI >= 0 && curJ >= 0 {
		pathNode := path[curI][curJ]
		fmt.Println(*pathNode)
		if curI-pathNode.prevI == 1 && curJ-pathNode.prevJ == 1 {
			pathList[ind] = &CurNode{i: curI, j: curJ}
			ind--
		}
		curI = pathNode.prevI
		curJ = pathNode.prevJ
	}

	return dp[rowSymCnt-1][colSymCnt-1], pathList[ind+1:]
}
