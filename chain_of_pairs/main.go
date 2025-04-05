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

	pairs := make([][2]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		pairs[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		pairs[i][1], _ = strconv.Atoi(scanner.Text())
	}

	list := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		list[i] = 1
	}
	for i := 1; i < cnt; i++ {
		for j := 0; j < i; j++ {
			if pairs[j][1] < pairs[i][0] && list[i] < list[j]+1 {
				list[i] = list[j] + 1
			}
		}
	}
	res := 1
	for _, v := range list {
		if res < v {
			res = v
		}
	}
	fmt.Println(res)
}
