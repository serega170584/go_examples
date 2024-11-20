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
	k, _ := strconv.Atoi(scanner.Text())

	list := make([]int, n)
	cnt := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	sum := make([]int, n)
	sum[0] = list[0]
	left := 0
	right := 0

	if sum[0] == k {
		cnt++
	}

	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + list[i]
		if sum[i] == k {
			cnt++
			right = i
			break
		}
		if sum[i] > k {
			right = i
			break
		}
	}

	if right == 0 {
		fmt.Println(cnt)
		return
	}

	for i := right; i < n; i++ {
		sum[i] = sum[i-1] + list[i]
		for sum[i]-sum[left] > k {
			left++
		}

		if sum[i]-sum[left] == k {
			cnt++
		}
	}

	fmt.Println(cnt)
}
