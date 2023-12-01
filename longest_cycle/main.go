package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

// RDRU
// [] R [] D [] R [] U []
// [0 0] [1 0] [1 1] [2 1] [2 0]
func main() {
	start := time.Now()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	list := make([]string, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		list[i] = scanner.Text()
	}

	locationNum := make(map[string]int, cnt+1)

	x, y := 0, 0
	ind := fmt.Sprintf("%d%d", x, y)
	num := 0
	locationNum[ind] = num
	max := 0

	for i := 0; i < cnt; i++ {
		num++

		way := list[i]

		if way == "L" {
			x -= 1
		}

		if way == "D" {
			y += 1
		}

		if way == "R" {
			x += 1
		}

		if way == "U" {
			y -= 1
		}

		ind = fmt.Sprintf("%d%d", x, y)

		if _, ok := locationNum[ind]; !ok {
			locationNum[ind] = num
		}

		cycleLen := num - locationNum[ind]

		if cycleLen > max {
			max = cycleLen
		}
	}

	fmt.Println("Max cycle length: ", max)

	fmt.Println("Estimated time: ", time.Since(start))
}
