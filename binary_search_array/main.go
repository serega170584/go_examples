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
	list := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Enter target:")
	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())

	l, r := 0, cnt-1
	for l <= r {
		middle := (l + r) / 2
		if list[middle] < target {
			l = middle + 1
			continue
		} else if list[middle] > target {
			r = middle - 1
			continue
		}
		fmt.Println(middle)
		return
	}

	fmt.Println(-1)
}
