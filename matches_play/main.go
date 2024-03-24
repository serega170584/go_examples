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
	n, _ := strconv.Atoi(scanner.Text())

	leftMatches := make([][4]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		leftMatches[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		leftMatches[i][1], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		leftMatches[i][2], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		leftMatches[i][3], _ = strconv.Atoi(scanner.Text())
	}

	rightMatches := make([][4]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		rightMatches[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		rightMatches[i][1], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		rightMatches[i][2], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		rightMatches[i][3], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(getMatchesMInCnt(n, leftMatches, rightMatches))
}

func getMatchesMInCnt(n int, leftMatches [][4]int, rightMatches [][4]int) int {
	rightVectors := make(map[[2]int][][2]int, n)
	for _, v := range rightMatches {
		vector := [2]int{v[2] - v[0], v[3] - v[1]}
		againstVector := [2]int{v[0] - v[2], v[1] - v[3]}
		vectorPoints := [2]int{v[0], v[1]}
		if _, ok := rightVectors[againstVector]; ok {
			vector = againstVector
			vectorPoints = [2]int{v[2], v[3]}
		}

		rightVectors[vector] = append(rightVectors[vector], vectorPoints)
	}

	diffList := make(map[[2]int]int, n)
	minCnt := n
	for _, v := range leftMatches {
		vector := [2]int{v[2] - v[0], v[3] - v[1]}
		againstVector := [2]int{v[0] - v[2], v[1] - v[3]}
		vectorPoints := [2]int{v[0], v[1]}
		if _, ok := rightVectors[againstVector]; ok {
			vector = againstVector
			vectorPoints = [2]int{v[2], v[3]}
		}

		if list, ok := rightVectors[vector]; ok {
			for _, rv := range list {
				diff := [2]int{rv[0] - vectorPoints[0], rv[1] - vectorPoints[1]}
				if _, diffOk := diffList[diff]; diffOk {
					diffList[diff]--
				} else {
					diffList[diff] = n - 1
				}
				if diffList[diff] < minCnt {
					minCnt = diffList[diff]
				}
			}
		}
	}
	return minCnt
}
