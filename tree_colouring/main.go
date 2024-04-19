package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// _ _ _ _ _ _ _
//     _ _ _ _
// min = 1
// max = 7
// max - min + 1 - (bl - tr + 1, bl >= tr) - (tl - br + 1, tl >= br) = 7

// _ _ _ _ _ _ _
//     _ _ _ _ _ _ _
// min = 1
// max = 9
// max - min + 1 - (bl - tr + 1, bl >= tr) - (tl - br + 1, tl >= br) = 9

//               _ _ _ _ _ _ _ _ _
// _ _ _ _ _ _ _
// 16 - 1 + 1 - (8 - 7 + 1) = 16

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
	topLeft := firstTree - maxFirstTrees
	topRight := firstTree + maxFirstTrees

	bottomLeft := secondTree - maxSecondTrees
	bottomRight := secondTree + maxSecondTrees

	minVal := min(topLeft, bottomLeft)
	maxVal := max(topRight, bottomRight)

	d := 0
	if topLeft > bottomRight {
		d += topLeft - bottomRight
	}

	if bottomLeft > topRight {
		d += bottomLeft - topRight
	}

	return maxVal - minVal + 1 - d
}
