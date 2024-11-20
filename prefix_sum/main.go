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

	list := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	sum := make([]int, cnt)
	sum[0] = list[0]

	for i := 1; i < cnt; i++ {
		sum[i] = sum[i-1] + list[i]
	}

	sum1 := make([]any, cnt)
	for i := 0; i < cnt; i++ {
		sum1[i] = sum[i]
	}

	fmt.Println(sum1...)
}
