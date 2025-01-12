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

	k := make([]int, n)
	list := make([][]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		k[i], _ = strconv.Atoi(scanner.Text())
		for j := 0; j < k[i]; j++ {
			scanner.Scan()
			list[i] = append(list[i], scanner.Text())
		}
	}

	playlistsDict := make(map[string]int, n)
	for _, row := range list {
		for _, v := range row {
			playlistsDict[v]++
		}
	}

	res := make([]any, 0)
	for s, v := range playlistsDict {
		if v == n {
			res = append(res, s)
		}
	}

	fmt.Println(len(res))
	fmt.Println(res...)
}
