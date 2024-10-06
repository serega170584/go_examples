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

	a := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	scanner.Scan()
	cnt, _ = strconv.Atoi(scanner.Text())

	b := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}

	scanner.Scan()
	cnt, _ = strconv.Atoi(scanner.Text())

	c := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		c[i], _ = strconv.Atoi(scanner.Text())
	}

	scanner.Scan()
	cnt, _ = strconv.Atoi(scanner.Text())
	d := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		d[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(convert(a, b, c, d))
}

func convert(list ...[]int) [][]int {
	ll := len(list)
	if len(list) == 0 {
		return [][]int{}
	}
	l := len(list[0])
	for i := 1; i < ll; i++ {
		il := len(list[i])
		if il < l {
			l = il
		}
	}
	res := make([][]int, 0, l)
	for i := 0; i < l; i++ {
		el := make([]int, 0, ll)
		for j := 0; j < ll; j++ {
			el = append(el, list[j][i])
		}
		res = append(res, el)
	}

	return res
}
