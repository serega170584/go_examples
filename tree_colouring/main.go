package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	firstTree, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	maxFirstTree, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	secondTree, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	maxSecondTree, _ := strconv.Atoi(scanner.Text())

	fmt.Println(maxColouringTrees(firstTree, maxFirstTree, secondTree, maxSecondTree))
}

func maxColouringTrees(firstTree int, maxFirstTrees int, secondTree int, maxSecondTrees int) int {
	leftFirstTree := firstTree - maxFirstTrees
	rightFirstTree := firstTree + maxFirstTrees

	leftSecondTree := secondTree - maxSecondTrees
	rightSecondTree := secondTree + maxSecondTrees

	points := make([]int, 4)
	points[0] = leftFirstTree
	points[1] = rightFirstTree
	points[2] = leftSecondTree
	points[3] = rightSecondTree
	slices.Sort(points)

	colouringTreesCnt := 0
	colouringTreesCnt += points[1] - points[0] + points[3] - points[2]

	if (points[2] == leftFirstTree && points[1] == rightSecondTree) || (points[2] == leftSecondTree && points[1] == rightFirstTree) {
		if points[1] != points[2] {
			colouringTreesCnt++
		}
	} else {
		colouringTreesCnt += points[2] - points[1]
	}

	colouringTreesCnt++

	return colouringTreesCnt
}
