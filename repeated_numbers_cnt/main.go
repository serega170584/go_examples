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
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	viMap := make(map[int]int, n)
	for i, v := range list {
		if prevI, ok := viMap[v]; ok {
			if i-prevI <= k {
				fmt.Println("YES")
				return
			}
		}

		viMap[v] = i
	}

	fmt.Println("NO")
}
