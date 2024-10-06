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

	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())

	values := make(map[int]struct{}, cnt)
	pairs := make([][2]int, 0, cnt)
	for _, v := range list {
		val := target - v
		if _, ok := values[val]; ok {
			pairs = append(pairs, [2]int{val, v})
		}
		values[v] = struct{}{}
	}

	fmt.Println(pairs)
}
