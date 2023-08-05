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

	rowLinks := make([]interface{}, colsCnt)
	weights := make([][]int, rowsCnt)
	for i := 0; i < rowsCnt; i++ {
		weights[i] = make([]int, colsCnt)
		for j := 0; j < colsCnt; j++ {
			rowLinks[j] = &weights[i][j]
		}
		_, err = fmt.Scanln(rowLinks...)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 1; i < rowsCnt; i++ {
		weights[i][0] += weights[i-1][0]
	}

	for j := 1; j < colsCnt; j++ {
		weights[0][j] += weights[0][j-1]
	}

	if rowsCnt == 1 || colsCnt == 1 {
		fmt.Println(weights[rowsCnt-1][colsCnt-1])
		return
	}

	for i := 1; i < rowsCnt; i++ {
		for j := 1; j < colsCnt; j++ {
			leftWeight := weights[i][j] + weights[i][j-1]
			upWeight := weights[i][j] + weights[i-1][j]
			if leftWeight < upWeight {
				weights[i][j] = leftWeight
			} else {
				weights[i][j] = upWeight
			}
		}
	}

	fmt.Println(weights[rowsCnt-1][colsCnt-1])
}
