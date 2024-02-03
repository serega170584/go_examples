package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	in := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		in[i], _ = strconv.Atoi(scanner.Text())
	}

	out := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		out[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Max visitors online: ", getMaxVisitorsOnline(n, in, out))

}

func getMaxVisitorsOnline(n int, in []int, out []int) int {
	maxOnline := 0
	online := 0
	events := make([][2]int, 2*n)
	for i := 0; i < n; i++ {
		events[2*i] = [2]int{in[i], -1}
		events[2*i+1] = [2]int{out[i], 1}
	}

	sort.Slice(events, func(i int, j int) bool {
		return events[i][0] < events[j][0]
	})

	for _, event := range events {
		if event[1] == -1 {
			online++
		} else {
			online--
		}
		maxOnline = max(online, maxOnline)
	}
	return maxOnline
}
