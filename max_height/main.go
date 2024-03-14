package main

import (
	"bufio"
	"fmt"
	"math"
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

	h, r := maxHeight(list, cnt)
	fmt.Println(h)
	fmt.Println(r)
}

func maxHeight(list [][2]int, cnt int) (int, string) {
	totalPositiveHeight := 0
	maxSecond := math.MinInt
	maxSecondIndex := 0
	negativeMaxUp := 0
	negativeMaxUpIndex := 0
	for i, v := range list {
		diff := v[0] - v[1]
		if diff >= 0 {
			totalPositiveHeight += diff
			if v[1] > maxSecond {
				maxSecond = v[1]
				maxSecondIndex = i
			}
		} else {
			if v[0] > negativeMaxUp {
				negativeMaxUp = v[0]
				negativeMaxUpIndex = i
			}
		}
	}

	cornerInd := 0
	cornerVal := 0
	if maxSecond > negativeMaxUp {
		cornerInd = maxSecondIndex
		cornerVal = maxSecond
	} else {
		cornerInd = negativeMaxUpIndex
		cornerVal = negativeMaxUp
	}
	maxHeightVal := totalPositiveHeight + cornerVal

	s := make([]string, cnt)
	left := 0
	right := cnt - 1
	for i, v := range list {
		if i != cornerInd {
			diff := v[0] - v[1]
			if diff < 0 {
				s[right] = strconv.Itoa(i + 1)
				right--
			} else {
				s[left] = strconv.Itoa(i + 1)
				left++
			}
		}
	}
	s[left] = strconv.Itoa(cornerInd + 1)

	return maxHeightVal, strings.Join(s, " ")
}
