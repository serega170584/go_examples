package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 6 - 2 = 4, 8 - 3 = 5
// 6 - 3 = 3, 8 - 2 = 6

// 0 1 0 0 0 0 0 0 0
// 0 1 1 0 0 0 0 0 0
// 0 0 0 0 0 0 1 1 0
// 0 0 0 0 0 1 0 0 0
// 0 0 0 1 0 0 0 0 0
// 0 0 0 0 0 0 0 0 0
// 0 0 0 0 0 0 0 0 1
// 0 0 0 0 0 0 0 0 0
// 1 0 0 0 0 0 0 0 0

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	lists := make([][]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		lists[i] = make([]int, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			lists[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}

	res := getMinCuts(cnt, lists)
	for _, v := range res {
		cnt = len(v)
		fmt.Println(cnt)
		s := make([]string, cnt)
		for i, val := range v {
			s[i] = strconv.Itoa(val)
		}
		fmt.Println(strings.Join(s, " "))
	}
}

func getMinCuts(cnt int, lists [][]int) [][]int {
	res := make([][]int, 0, cnt)
	for _, list := range lists {
		capacity := 200000
		sum := 0
		sums := make([]int, 0)
		counter := 0
		for i, v := range list {
			capacity = min(capacity, v)
			counter++
			if capacity < counter {
				sums = append(sums, sum)
				capacity = v
				sum = 0
				counter = 1
			}
			sum++

			if i == len(list)-1 {
				sums = append(sums, sum)
			}
		}
		res = append(res, sums)
	}

	return res
}
