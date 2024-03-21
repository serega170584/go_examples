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

	cnt := 0

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	a := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}
	cnt += n

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	b := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}
	cnt += n

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	c := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		c[i], _ = strconv.Atoi(scanner.Text())
	}
	cnt += n

	list := getSharedListsNumbers(cnt, a, b, c)
	sList := make([]string, 0, len(list))
	for _, v := range list {
		sList = append(sList, strconv.Itoa(v))
	}

	fmt.Println(strings.Join(sList, " "))
}

func getSharedListsNumbers(cnt int, a []int, b []int, c []int) []int {
	numbersCnts := make(map[int]int, cnt)

	aCnts := make(map[int]struct{}, cnt)
	for _, v := range a {
		aCnts[v] = struct{}{}
	}

	for i := range aCnts {
		numbersCnts[i]++
	}

	bCnts := make(map[int]struct{}, cnt)
	for _, v := range b {
		bCnts[v] = struct{}{}
	}

	for i := range bCnts {
		numbersCnts[i]++
	}

	cCnts := make(map[int]struct{}, cnt)
	for _, v := range c {
		cCnts[v] = struct{}{}
	}

	for i := range cCnts {
		numbersCnts[i]++
	}

	numbers := make([]int, 0, cnt)
	for n, v := range numbersCnts {
		if v >= 2 {
			numbers = append(numbers, n)
		}
	}

	slices.Sort(numbers)

	return numbers
}
