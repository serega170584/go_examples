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

	fmt.Println("Enter events count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter in events")
	in := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		in[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Enter out events")
	out := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		out[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Enter boss events count")
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter boss events")
	bossEvents := make([]int, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		bossEvents[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(getBossCounters(n, in, out, m, bossEvents))
}

func getBossCounters(n int, in []int, out []int, m int, boss []int) []int {
	bossCounters := make([]int, m)
	online := 0
	events := make([][2]int, 2*n+m)
	for i := 0; i < n; i++ {
		events[2*i] = [2]int{in[i], -1}
		events[2*i+1] = [2]int{out[i], 1}
	}

	for i := 0; i < m; i++ {
		events[2*n+i] = [2]int{boss[i], 0}
	}

	sort.Slice(events, func(i int, j int) bool {
		return events[i][0] < events[j][0] || (events[i][0] == events[j][0] && events[i][1] < events[j][1])
	})

	bossCounterIndex := 0
	for _, event := range events {
		if event[1] == -1 {
			online++
		} else if event[1] == 1 {
			online--
		} else if event[1] == 0 {
			bossCounters[bossCounterIndex] = online
			bossCounterIndex++
		}
	}
	return bossCounters
}
