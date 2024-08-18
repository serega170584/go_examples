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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	b := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	c := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		c[i], _ = strconv.Atoi(scanner.Text())
	}

	list := getSharedListsNumbers(a, b, c)
	sList := make([]string, 0, len(list))
	for _, v := range list {
		sList = append(sList, strconv.Itoa(v))
	}

	fmt.Println(strings.Join(sList, " "))
}

func getSharedListsNumbers(a []int, b []int, c []int) []int {
	m := make(map[int]int)

	am := make(map[int]struct{})
	for _, v := range a {
		if _, ok := am[v]; !ok {
			am[v] = struct{}{}
			m[v]++
		}
	}

	bm := make(map[int]struct{})
	for _, v := range b {
		if _, ok := bm[v]; !ok {
			bm[v] = struct{}{}
			m[v]++
		}
	}

	cm := make(map[int]struct{})
	for _, v := range c {
		if _, ok := cm[v]; !ok {
			cm[v] = struct{}{}
			m[v]++
		}
	}

	l := make([]int, 0)
	for i, v := range m {
		if v >= 2 {
			l = append(l, i)
		}
	}

	slices.Sort(l)

	return l
}
