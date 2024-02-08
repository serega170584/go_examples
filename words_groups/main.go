package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter list count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list")
	list := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i] = scanner.Text()
	}

	fmt.Println(getWordsGroups(n, list))
}

func getWordsGroups(n int, s []string) [][]string {
	groups := make([][]string, 0)
	groupsMap := make(map[string][]string, n)

	for _, val := range s {
		r := []rune(val)
		sort.Slice(r, func(i int, j int) bool {
			return r[i] < r[j]
		})
		s := string(r)
		if _, ok := groupsMap[s]; !ok {
			groupsMap[s] = make([]string, 1)
		}
		groupsMap[s] = append(groupsMap[s], val)
	}

	for _, v := range groupsMap {
		groups = append(groups, v)
	}

	return groups
}
