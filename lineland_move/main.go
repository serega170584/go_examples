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
	n, _ := strconv.Atoi(scanner.Text())

	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	res := make([]any, n)
	res[n-1] = -1

	q := make([]int, 0, n)
	q = append(q, n-1)

	minVal := list[n-1]
	for i := n - 2; i >= 0; i-- {
		if list[i+1] == list[i] {
			res[i] = res[i+1]
			q = append(q, i)
			continue
		}
		if list[i+1] > list[i] && res[i+1] == -1 {
			minVal = list[i]
			res[i] = res[i+1]
			q = append(q, i)
			continue
		}
		if list[i] <= minVal {
			minVal = list[i]
			res[i] = -1
			q = append(q, i)
			continue
		}
		ql := len(q)
		for j := len(q) - 1; j >= 0; j-- {
			ind := q[j]
			if list[ind] < list[i] {
				res[i] = ind
				q = append(q, i)
				break
			}
		}
		if ql == len(q) {
			res[i] = -1
			q = append(q, i)
		}
	}

	fmt.Println(res...)
}
