package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	list := make([][2]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		list[i][0], _ = strconv.Atoi(scanner.Text())

		scanner.Scan()
		list[i][1], _ = strconv.Atoi(scanner.Text())
	}

	path, height := getPathHeight(list)

	strList := make([]string, len(path))
	for i := range strList {
		strList[i] = strconv.Itoa(path[i] + 1)
	}

	fmt.Println(height, strings.Join(strList, " "))
}

func getPathHeight(list [][2]int) ([]int, int) {
	minPositiveDiff := 0
	maxItemHeight := 0
	minPositiveIndex := -1
	positiveIndexList := make([]int, 0)
	positiveSum := 0
	maxNegativeItemHeight := 0
	maxNegativeIndex := -1
	negativeIndexList := make([]int, 0)
	for i, v := range list {
		if v[0] > v[1] {
			positiveSum += v[0] - v[1]
			if v[0] > maxItemHeight {
				maxItemHeight = v[0]
				minPositiveDiff = v[0] - v[1]
				minPositiveIndex = len(positiveIndexList)
			} else if v[0] == maxItemHeight {
				if v[0]-v[1] < minPositiveDiff {
					minPositiveDiff = v[0] - v[1]
					minPositiveIndex = len(positiveIndexList)
				}
			}
			positiveIndexList = append(positiveIndexList, i)
		} else {
			if v[0] > maxNegativeItemHeight {
				maxNegativeItemHeight = v[0]
				maxNegativeIndex = len(negativeIndexList)
			}
			negativeIndexList = append(negativeIndexList, i)
		}
	}

	height := positiveSum
	positiveSum -= minPositiveDiff
	positiveSum += maxItemHeight

	if minPositiveIndex != -1 {
		positiveIndexList[minPositiveIndex], positiveIndexList[len(positiveIndexList)-1] = positiveIndexList[len(positiveIndexList)-1], positiveIndexList[minPositiveIndex]
	}

	if maxNegativeIndex != -1 {
		negativeIndexList[maxNegativeIndex], negativeIndexList[0] = negativeIndexList[0], negativeIndexList[maxNegativeIndex]
	}

	return append(positiveIndexList, negativeIndexList...), max(positiveSum, height+maxNegativeItemHeight)

}
