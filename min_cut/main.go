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
	n, _ := strconv.Atoi(scanner.Text())

	sets := make([][]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		l, _ := strconv.Atoi(scanner.Text())
		for j := 0; j < l; j++ {
			scanner.Scan()
			v, _ := strconv.Atoi(scanner.Text())
			sets[i] = append(sets[i], v)
		}
	}

	res := make([][]int, n)
	for i, set := range sets {
		minVal := set[0]
		cutLen := 1
		for j := 1; j < len(set); j++ {
			if set[j] < minVal {
				minVal = set[j]
			}
			cutLen++
			if minVal < cutLen {
				cutLen--
				res[i] = append(res[i], cutLen)
				cutLen = 1
				minVal = set[j]
				continue
			}
		}
		res[i] = append(res[i], cutLen)
	}

	for _, row := range res {
		cur := make([]any, len(row))
		for i, v := range row {
			cur[i] = v
		}
		fmt.Println(len(row))
		fmt.Println(cur...)
	}
}
