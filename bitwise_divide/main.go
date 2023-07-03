package main

import "fmt"

func main() {
	fmt.Println(divide(10, 3))
}

func divide(num int, divisor int) int {
	var result int
	for num > divisor {
		middleResult := 1
		lastOffset := divisor
		offset := divisor << 1
		for offset < num {
			middleResult = middleResult << 1
			lastOffset = offset
			offset = lastOffset << 1
		}
		result += middleResult
		num -= lastOffset
	}
	return result
}
