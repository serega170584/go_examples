package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	aInd := a / k
	if a%k != 0 {
		aInd++
	}

	bInd := b / k
	if b%k != 0 {
		bInd++
	}

	diff := aInd - bInd
	if diff < 0 {
		diff = -diff
	}

	startInd := 0
	endInd := n - 1
	otherWayStartIndex := 0
	otherWayEndIndex := n - 1

	if diff < n {
		mod := aInd % n
		if mod == 0 {
			startInd = n
			otherWayStartIndex = 2
		} else {
			startInd = mod
			if mod == 1 {
				otherWayStartIndex = 1
			} else {
				otherWayStartIndex = n - mod + 2
			}
		}
		startInd--
		otherWayStartIndex--

		mod = bInd % n
		if mod == 0 {
			endInd = n
			otherWayEndIndex = 1
		} else {
			endInd = mod
			if mod == 1 {
				otherWayEndIndex = 1
			} else {
				otherWayEndIndex = n - mod + 2
			}
		}
		endInd--
		otherWayEndIndex--

		if endInd-startInd < 0 {
			startInd, endInd = endInd, startInd
		}

		if otherWayEndIndex-otherWayStartIndex < 0 {
			otherWayStartIndex, otherWayEndIndex = otherWayEndIndex, otherWayStartIndex
		}
	}

	maxWon := math.MinInt
	for i := startInd; i <= endInd; i++ {
		maxWon = max(maxWon, list[i])
	}

	for i := otherWayStartIndex; i <= otherWayEndIndex; i++ {
		maxWon = max(maxWon, list[i])
	}

	fmt.Println(maxWon)
}
