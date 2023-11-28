package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

// 6
// 1 2 3 4 5 6
// 6 + 1 = 7
// 7 / 2 = 3
// 3 * 3 = 9 > 6
// 4 * 4 = 16 > 6
// 1 2 3
// 1 + 3 = 4
// 4 / 2 = 2
// 2 * 2 = 4 < 6
// 3 * 3 = 9 > 6
// solution = 2
func main() {
	t := time.Now()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	res := sqrt(x)
	fmt.Println(res)
	fmt.Println("Elapsed: ", time.Since(t))
}

func sqrt(x int) int {
	left, right := 1, x
	var middle int
	for left <= right {
		middle = (left + right) / 2
		if middle*middle <= x && (middle+1)*(middle+1) > x {
			break
		} else if middle*middle < x {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return middle
}
