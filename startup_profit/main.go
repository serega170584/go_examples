package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	d, _ := strconv.Atoi(scanner.Text())

	fmt.Println(getMaxStratupProfit(n, k, d))
}

func getMaxStratupProfit(n int, k int, d int) string {
	m := 0

	res := make([]int, 0)
	delim := n
	for delim != 0 {
		res = append([]int{delim % 10}, res...)
		delim /= 10
	}

	for i := 0; i < d; i++ {
		for j := 0; j < 10; j++ {
			m = (n + m) % k
		}

		v := -1

		for j := 0; j < 10; j++ {
			n1 := m + j
			if n1%k == 0 {
				v = j
				n = n1
			}
		}

		if v == -1 {
			return "-1"
		}

		res = append(res, v)
	}

	var sb strings.Builder

	for _, v := range res {
		sb.WriteString(strconv.Itoa(v))
	}

	return sb.String()
}
