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

	firstSum := make([]int, n)
	firstSum[0] = 0
	for i := 1; i < n; i++ {
		firstSum[i] = (firstSum[i-1] + list[i-1]) % 1000000007
	}

	// 1000000006 1000000
	firstMulti := make([]int, n)
	for i := 0; i < n; i++ {
		delim := firstSum[i] / 2000
		mod := firstSum[i] % 2000
		firstMulti[i] = ((((list[i]*2000)%1000000007)*delim)%1000000007 + (list[i]*mod)%1000000007) % 1000000007
	}

	secondSum := make([]int, n)
	secondSum[0] = 0
	secondSum[1] = 0
	for i := 2; i < n; i++ {
		secondSum[i] = (secondSum[i-1] + firstMulti[i-1]) % 1000000007
	}

	secondMulti := make([]int, n)
	for i := 2; i < n; i++ {
		delim := secondSum[i] / 2000
		mod := secondSum[i] % 2000
		secondMulti[i] = (secondMulti[i-1] + ((((list[i]*2000)%1000000007)*delim)%1000000007+(list[i]*mod)%1000000007)%1000000007) % 1000000007
	}

	fmt.Println(secondMulti[n-1])
}
