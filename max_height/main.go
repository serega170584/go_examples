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

	list := make([][2]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		list[i][1], _ = strconv.Atoi(scanner.Text())
	}

	positiveList := make([]int, 0, n)
	negativeList := make([]int, 0, n)
	maxPositiveDown := math.MinInt
	maxNegativeUp := math.MinInt
	maxPositiveListIndex := -1
	maxNegativeListIndex := -1
	weight := 0
	for i, v := range list {
		if v[0]-v[1] > 0 {
			positiveList = append(positiveList, i)
			weight += v[0] - v[1]
			if v[1] > maxPositiveDown {
				maxPositiveDown = v[1]
				maxPositiveListIndex = len(positiveList) - 1
			}
		} else {
			negativeList = append(negativeList, i)
			if v[0] > maxNegativeUp {
				maxNegativeUp = v[0]
				maxNegativeListIndex = len(negativeList) - 1
			}
		}
	}

	if maxPositiveListIndex != -1 {
		positiveList[len(positiveList)-1], positiveList[maxPositiveListIndex] = positiveList[maxPositiveListIndex], positiveList[len(positiveList)-1]
	}
	if maxNegativeListIndex != -1 {
		negativeList[0], negativeList[maxNegativeListIndex] = negativeList[maxNegativeListIndex], negativeList[0]
	}
	weight = max(weight+maxPositiveDown, weight+maxNegativeUp)

	res := append(positiveList, negativeList...)
	for i := range res {
		res[i]++
	}

	fmt.Println(res)
}
