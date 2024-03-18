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
	m, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	text := make([][]rune, m)
	for i := 0; i < m; i++ {
		text[i] = make([]rune, n)
		scanner.Scan()
		text[i] = []rune(scanner.Text())
	}

	t, res := isIsolatedRects(text)
	if t {
		fmt.Println("YES")
		fmt.Println(res)
		return
	}

	fmt.Println("NO")
}

func isIsolatedRects(text [][]rune) (bool, string) {
	lattice := []rune("#")[0]

	minVal1 := math.MaxInt
	maxVal1 := math.MinInt
	rowFrom1 := math.MinInt
	rowTo1 := math.MinInt
	isMax1Defined := false
	minVal2 := math.MaxInt
	maxVal2 := math.MinInt
	rowFrom2 := math.MinInt
	rowTo2 := math.MinInt
	isMax2Defined := false

	for i, row := range text {
		curMinVal := math.MaxInt
		curMaxVal := math.MinInt
		for j, v := range row {
			if v == lattice {
				if j > curMaxVal+1 && curMaxVal != math.MinInt {
					if !isMax1Defined {
						isMax1Defined = true
						rowFrom1 = i
						rowTo1 = i
					} else {
						rowTo1++
					}

					minVal1 = curMinVal
					maxVal1 = curMaxVal
					curMinVal = math.MaxInt
					curMaxVal = math.MinInt
				}

				curMinVal = min(curMinVal, j)
				curMaxVal = max(curMaxVal, j)
			}
		}

		if curMinVal == math.MaxInt && curMaxVal == math.MinInt {
			continue
		}

		if isMax1Defined && isMax2Defined {
			if i > rowTo1+1 && i > rowTo2+1 {
				return false, ""
			}
		}

		if !isMax1Defined || (minVal1 == curMinVal && maxVal1 == curMaxVal) {
			if minVal1 == curMinVal && maxVal1 == curMaxVal && isMax2Defined && i-rowTo1 > 1 && i-rowTo2 > 1 {
				return false, ""
			}

			if minVal1 == curMinVal && maxVal1 == curMaxVal && !isMax2Defined && i-rowTo1 > 1 {
				minVal2 = curMinVal
				maxVal2 = curMaxVal
				rowFrom2, rowTo2 = i, i
				isMax2Defined = true
			}

			if rowTo2 == i-1 {
				rowTo1 = i
				if minVal1 > minVal2 {
					maxVal2 = minVal1 - 1
				} else {
					minVal2 = maxVal1 + 1
				}
			} else {
				minVal1 = curMinVal
				maxVal1 = curMaxVal
				if !isMax1Defined {
					isMax1Defined = true
					rowFrom1 = i
					rowTo1 = i
				} else {
					rowTo1++
				}
			}
		} else if !isMax2Defined || (minVal2 == curMinVal && maxVal2 == curMaxVal) {
			minVal2 = curMinVal
			maxVal2 = curMaxVal
			if !isMax2Defined {
				isMax2Defined = true
				rowFrom2 = i
				rowTo2 = i
			} else {
				rowTo2++
			}
		} else {
			newMin := 0
			if minVal1 == minVal2 {
				newMin = maxVal1 + 1
				if newMin == curMinVal {
					minVal2 = curMinVal
					rowTo1 += rowFrom2 - rowFrom1 + 1
					rowTo2++
				} else {
					return false, ""
				}
			}

			if maxVal1 == maxVal2 {
				newMin = minVal1 - 1
				if newMin == curMaxVal {
					rowTo1 += rowFrom2 - rowFrom1 + 1
					rowTo2++
					minVal2 = curMinVal
					maxVal2 = curMaxVal
				} else {
					return false, ""
				}
			}
		}
	}

	if !isMax2Defined {
		if rowFrom1 == rowTo1 {
			if maxVal1 > minVal1 {
				minVal2, maxVal2 = minVal1, minVal1
				rowFrom2, rowTo2 = rowFrom1, rowTo1
				minVal1++
			} else {
				return false, ""
			}
		} else {
			minVal2, maxVal2 = minVal1, maxVal1
			rowFrom2, rowTo2 = rowFrom1, rowFrom1
			rowFrom1++
		}
	}

	var sb strings.Builder
	for i, row := range text {
		for j, v := range row {
			if v == lattice {
				if rowFrom1 <= i && i <= rowTo1 && minVal1 <= j && j <= maxVal1 {
					sb.WriteString("a")
				} else if rowFrom2 <= i && i <= rowTo2 && minVal2 <= j && j <= maxVal2 {
					sb.WriteString("b")
				}
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}
	return true, sb.String()
}
