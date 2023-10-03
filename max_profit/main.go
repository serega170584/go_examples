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
	cnt, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	var min, maxProfit int
	min = int(math.Inf(1))
	maxProfit = int(math.Inf(-1))

	for _, val := range arr {
		if val < min {
			min = val
		}

		profit := val - min
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	fmt.Println(maxProfit)

}
