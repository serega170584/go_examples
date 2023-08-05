package main

import (
	"fmt"
	"log"
)

func main() {
	var rowsCnt, colsCnt int
	_, err := fmt.Scanln(&rowsCnt, &colsCnt)
	if err != nil {
		log.Fatal(err)
	}

	weights := make([][]int, rowsCnt)
	filledWeights := make([][]int, rowsCnt)

	row := make([]interface{}, colsCnt)

	for i := 0; i < rowsCnt; i++ {
		weights[i] = make([]int, colsCnt)
		filledWeights[i] = make([]int, colsCnt)
		for j := 0; j < colsCnt; j++ {
			row[j] = &weights[i][j]
			filledWeights[i][j] = -1
		}

		_, err = fmt.Scanln(row...)
		if err != nil {
			log.Fatal(err)
		}
	}

	filledWeights = longestRoute(weights, filledWeights, rowsCnt-1, colsCnt-1)

	fmt.Println(filledWeights)
}

func longestRoute(weights, filledWeights [][]int, rowInd, colInd int) [][]int {
	var weight, leftWeight, upWeight int

	if filledWeights[rowInd][colInd] != -1 {
		return filledWeights
	}

	if rowInd == 0 && colInd == 0 {
		filledWeights[rowInd][colInd] = weights[rowInd][colInd]
		return filledWeights
	}

	if colInd == 0 {
		leftWeight = -1
	} else {
		filledWeights = longestRoute(weights, filledWeights, rowInd, colInd-1)
		leftWeight = filledWeights[rowInd][colInd-1]
	}

	if rowInd == 0 {
		upWeight = -1
	} else {
		filledWeights = longestRoute(weights, filledWeights, rowInd-1, colInd)
		upWeight = filledWeights[rowInd-1][colInd]
	}

	weight = leftWeight
	if upWeight > leftWeight {
		weight = upWeight
	}

	filledWeights[rowInd][colInd] = weights[rowInd][colInd] + weight

	return filledWeights
}
