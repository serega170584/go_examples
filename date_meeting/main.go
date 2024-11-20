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

	scanner.Scan()
	r, _ := strconv.Atoi(scanner.Text())

	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	interval := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		interval[i] = list[i+1] - list[i]
	}

	sum := make([]int, n-1)
	sum[0] = interval[0]
	for i := 1; i < n-1; i++ {
		sum[i] = sum[i-1] + interval[i]
	}

	left := -1
	right := -1
	for i := 0; i < n-1; i++ {
		if sum[i] > r {
			right = i
			break
		}
	}

	cnt := 0

	if right == -1 {
		fmt.Println(cnt)
		return
	}

	if right == 0 {
		right++
		cnt++
	}

	for i := right; i < n-1; i++ {
		cnt++
		cnt += left + 1
		for sum[i]-sum[left+1] > r {
			left++
			cnt++
		}
	}

	fmt.Println(cnt)

}
