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

	arr := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	for i := 0; i < cnt; i++ {
		seen := make(map[int]struct{}, cnt)

		for j := i; j < cnt; j++ {
			if i != j {
				negSum := -arr[i] - arr[j]
				if _, ok := seen[negSum]; ok {
					fmt.Println(arr[i], arr[j], negSum)
				}
				seen[arr[j]] = struct{}{}
			}
		}
	}
}
