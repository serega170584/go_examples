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
	l, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	x1, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	v1, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	x2, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	v2, _ := strconv.Atoi(scanner.Text())

	timeVal := getTime(l, x1, v1, x2, v2)
	if timeVal == math.MaxFloat64 {
		fmt.Println("NO")
		return
	}

	fmt.Println("YES")
	fmt.Println(fmt.Sprintf("%.10f", timeVal))
}

func getTime(l int, x1 int, v1 int, x2 int, v2 int) float64 {
	minVal := math.MaxFloat64

	if x1 == x2 {
		return 0
	}

	if l-x1 == x2 {
		return 0
	}

	if v2-v1 == 0 && v2+v1 == 0 {
		return minVal
	}

	if x1 < 0 {
		x1 = l + x1
	}

	if x2 < 0 {
		x2 = l + x2
	}

	if v2-v1 != 0 {
		t := float64(x1-x2) / float64(v2-v1)
		if t >= 0 {
			minVal = t
		}
	}

	if v1+v2 != 0 {
		t := float64(l-x1-x2) / float64(v1+v2)
		if t >= 0 {
			minVal = min(minVal, t)
		}
	}

	if minVal != math.MaxFloat64 {
		return minVal
	}

	minCornerTime := math.MaxFloat64
	cornerTime := -float64(x1) / float64(v1)
	newX1 := float64(0)
	newX2 := float64(0)
	if cornerTime >= 0 {
		minCornerTime = cornerTime
		newX1 = float64(l)
		newX2 = float64(x2) + float64(v2)*cornerTime
	}

	cornerTime = float64(l-x1) / float64(v1)
	if cornerTime >= 0 {
		minCornerTime = min(minCornerTime, cornerTime)
		if minCornerTime == cornerTime {
			newX1 = float64(0)
			newX2 = float64(x2) + float64(v2)*cornerTime
		}
	}

	cornerTime = -float64(x2) / float64(v2)
	if cornerTime >= 0 {
		minCornerTime = min(minCornerTime, cornerTime)
		if minCornerTime == cornerTime {
			newX2 = float64(l)
			newX1 = float64(x1) + float64(v1)*cornerTime
		}
	}

	cornerTime = float64(l-x2) / float64(v2)
	if cornerTime >= 0 {
		minCornerTime = min(minCornerTime, cornerTime)
		if minCornerTime == cornerTime {
			newX2 = float64(0)
			newX1 = float64(x1) + float64(v1)*cornerTime
		}
	}

	if v2-v1 != 0 {
		t := (newX1 - newX2) / float64(v2-v1)
		if t >= 0 {
			minVal = t
		} else {
			minVal = -t
		}
	}

	if v1+v2 != 0 {
		t := (float64(l) - newX1 - newX2) / float64(v1+v2)
		if t >= 0 {
			minVal = min(minVal, t)
		} else {
			t = -t
			minVal = min(minVal, t)
		}
	}

	minVal += minCornerTime
	return minVal
}
