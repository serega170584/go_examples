package main

import (
	"fmt"
	"math"
)

func main() {
	tests := make([][5]int, 0)
	//tests = append(tests, [5]int{100, 96, 2, 97, 1})
	//tests = append(tests, [5]int{100, 97, 1, 96, 2})
	//tests = append(tests, [5]int{100, 4, -1, 94, 3})
	//tests = append(tests, [5]int{100, 94, 3, 4, -1})
	//tests = append(tests, [5]int{100, 86, 1, 94, -3})
	//tests = append(tests, [5]int{100, 94, -3, 86, 1})
	//tests = append(tests, [5]int{100, 86, 5, 98, 2})
	//tests = append(tests, [5]int{100, 98, 2, 86, 5})
	//tests = append(tests, [5]int{100, -86, -5, -98, -2})
	//tests = append(tests, [5]int{100, -98, -2, -86, -5})
	//tests = append(tests, [5]int{6, 3, 1, 1, 1})
	tests = append(tests, [5]int{12, 8, 10, 5, 20})
	//tests = append(tests, [5]int{5, 0, 0, 1, 2})
	//tests = append(tests, [5]int{10, 7, -3, 1, 4})
	for _, test := range tests {
		fmt.Println(test)
		fmt.Println("Result: ", getTime(test[0], test[1], test[2], test[3], test[4]))
	}
}

// 100
// 100 = 60 +40
// -60 40
// x1 + v1 * t = l - x2 - v2 * t,  t = (l - x1 - x2) / (v1 + v2)

// (100 + 60 - 40)

// x1 + v1 * t = x2 + v2 * t => t = (x1 - x2) / (v2 - v1)
// x1 + v1 * t = l - x2 - v2 * t =>  t = (l - x2 -)
// 0 - x1 - v1 * t = x2 + v2 * t
//
// 100, 86, 5, 98, 2
// -l -x1 - v1 * t = x2 + v2 * t
// t = (-l - x1 - x2) / (v1+v2)
// -l + x1 + v1 * t =

func getTime(l int, x1 int, v1 int, x2 int, v2 int) float64 {
	minVal := math.MaxFloat64

	if v2-v1 == 0 && v2+v1 == 0 {
		return minVal
	}

	if v2+v1 == 0 {
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

	if v2-v1 > 0 {
		t := (newX1 - newX2) / float64(v2-v1)
		if t >= 0 {
			minVal = t
		}
	}

	if v1+v2 > 0 {
		t := (float64(l) - newX1 - newX2) / float64(v1+v2)
		if t >= 0 {
			minVal = min(minVal, t)
		}
	}

	minVal += cornerTime
	return minVal
}
