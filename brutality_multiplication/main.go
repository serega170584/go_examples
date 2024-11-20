package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// aacccdddeebddeebaqwertybppppppbb
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	c := scanner.Text()
	ic := [2]int{}
	i := len(c) - 9
	first := 0
	if i >= 0 {
		first = i
	}
	ic[1], _ = strconv.Atoi(c[first:])
	ic[0], _ = strconv.Atoi(c[:first])

	scanner.Scan()
	s := scanner.Text()

	aCnt := 0
	cnt := 1
	maxCnt := cnt
	if s[0] == []byte("a")[0] {
		aCnt++
	}
	sum := make([][2]int, n)
	sum[0] = [2]int{0, 1}
	start := 0
	for i = 1; i < n; i++ {
		cnt++
		if s[i] == []byte("a")[0] {
			aCnt++
		}
		extra := 0
		tailSum := sum[i-1][1]
		if s[i] == []byte("b")[0] {
			tailSum += aCnt
		}
		if tailSum > 1000000000 {
			extra = tailSum / 1000000000
			tailSum %= 1000000000
		}
		sum[i] = [2]int{sum[i-1][0] + extra, tailSum}
		diff := [2]int{}
		right := sum[i][1]
		diffExtra := 0
		if sum[i][1] < sum[start][1] {
			right += 1000000000
			diffExtra = 1
		}
		diff[1] = right - sum[start][1]
		diff[0] = sum[i][0] - sum[start][0] - diffExtra
		for diff[1] > ic[1] || diff[0] > ic[0] {
			cnt--
			start++
			diffExtra = 0
			if sum[i][1] < sum[start][1] {
				right += 1000000000
				diffExtra = 1
			}
			diff[1] = right - sum[start][1] - aCnt
			diff[0] = sum[i][0] - sum[start][0] - diffExtra
		}
		if cnt > maxCnt {
			maxCnt = cnt
		}
	}

	fmt.Println(maxCnt)
}
