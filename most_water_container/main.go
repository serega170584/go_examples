package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	container := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		container[i], _ = strconv.Atoi(scanner.Text())
	}

	left := 0
	right := cnt - 1

	maxSquare := 0
	for left != right {
		var height, weight int
		weight = right - left
		if container[left] > container[right] {
			height = right
			right--
		} else {
			height = left
			left++
		}
		square := height * weight
		if square > maxSquare {
			maxSquare = square
		}
	}

	fmt.Println(maxSquare)
}
