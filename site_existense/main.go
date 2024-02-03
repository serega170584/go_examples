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

	fmt.Println("Time with visitors: ", timeWithVisitors(n, in, out))
}

func timeWithVisitors(n int, in []int, out []int) int {
	online := 0
	timeWithVisitors := 0
	events := make([][2]int, 2*n)
	for i := 0; i < n; i++ {
		events[2*i] = [2]int{in[i], -1}
		events[2*i+1] = [2]int{out[i], 1}
	}

	sort.Slice(events, func(i int, j int) bool {
		return events[i][0] < events[j][0]
	})

	for i, event := range events {
		if online > 0 {
			timeWithVisitors += events[i][0] - events[i-1][0]
		}
		if event[1] == -1 {
			online += 1
		} else {
			online -= 1
		}
	}
	return timeWithVisitors
}
