package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode/utf8"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter string")
	scanner.Scan()
	s := scanner.Text()

	printChart(s)
}

func printChart(s string) {
	dict := make(map[rune]int, utf8.RuneCountInString(s))

	maxCount := 0
	uniqueSymCnt := 0
	for _, v := range s {
		if _, ok := dict[v]; !ok {
			dict[v] = 0
			uniqueSymCnt++
		}
		dict[v]++
		maxCount = max(maxCount, dict[v])
	}

	list := make([]rune, uniqueSymCnt)
	listInd := 0
	for r, _ := range dict {
		list[listInd] = r
		listInd++
	}

	sort.Slice(list, func(i int, j int) bool {
		return list[i] < list[j]
	})

	for i := maxCount; i > 0; i-- {
		for _, v := range list {
			if dict[v] >= i {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
