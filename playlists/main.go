package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	items := make([][]string, cnt)

	for i := 0; i < cnt; i++ {
		scanner.Scan()
		itemsCnt, _ := strconv.Atoi(scanner.Text())
		for j := 0; j < itemsCnt; j++ {
			scanner.Scan()
			items[i] = append(items[i], scanner.Text())
		}
	}

	m := make(map[string]int)
	for _, item := range items {
		for _, v := range item {
			m[v]++
		}
	}

	list := make([]string, 0)
	for track, trackCnt := range m {
		if trackCnt == cnt {
			list = append(list, track)
		}
	}
	slices.Sort(list)

	printList := make([]interface{}, len(list))
	for i, v := range list {
		printList[i] = v
	}

	fmt.Println(printList...)
}
