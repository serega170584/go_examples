package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	k := 0
	lists := make([][]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		t, _ := strconv.Atoi(scanner.Text())

		lists[i] = make([]string, t)
		k += t
		for j := 0; j < t; j++ {
			scanner.Scan()
			lists[i][j] = scanner.Text()
		}
	}

	l, c := getPlaylists(n, k, lists)

	fmt.Println(l)
	fmt.Println(strings.Join(c, " "))
}

func getPlaylists(n int, k int, lists [][]string) (int, []string) {
	m := make(map[string]int, n)
	c := make([]string, 0, n)

	for _, list := range lists {
		for _, v := range list {
			m[v]++
			if m[v] == n {
				c = append(c, v)
			}
		}
	}

	slices.Sort(c)

	return len(c), c
}
