package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var list []int
	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		list = append(list, v)
	}

	m := make(map[int]int, 4)
	m[0] = 0
	m[1] = 1
	m[2] = 2
	m[3] = 2

	cnt := 0
	for _, v := range list {
		cnt += v/4 + m[v%4]
	}

	fmt.Println(cnt)
}
