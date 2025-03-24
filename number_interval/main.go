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
	searchCnt, _ := strconv.Atoi(scanner.Text())

	searchList := make([][2]int, searchCnt)
	for i := 0; i < searchCnt; i++ {
		scanner.Scan()
		searchList[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		searchList[i][1], _ = strconv.Atoi(scanner.Text())
	}

	res := make([]int, searchCnt)
	for i := 0; i < searchCnt; i++ {
		res[i] = rightSearch(list, searchList[i][1]) - leftSearch(list, searchList[i][0])
		if res[i] >= 0 {
			res[i] = res[i] + 1
			continue
		}
		res[i] = 0
	}

	fmt.Println(res)
}

func leftSearch(a []int, v int) int {
	l := 0
	r := len(a) - 1
	for l < r {
		m := (l + r) / 2
		if a[m] >= v {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func rightSearch(a []int, v int) int {
	l := 0
	r := len(a) - 1
	for l < r {
		m := (l + r + 1) / 2
		if a[m] <= v {
			l = m
		} else {
			r = m - 1
		}
	}

	return l
}
