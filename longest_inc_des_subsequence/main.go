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

	a := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(longestSubsequenceLength(a))

}

func longestSubsequenceLength(a []int) int {
	if len(a) < 3 {
		return 0
	}

	cnt := len(a)

	longestIncLength := make([]int, cnt)
	longestIncLength[0] = 1
	for i := 1; i < cnt; i++ {
		longestIncLength[i] = 1
		for j := 0; j < i; j++ {
			if a[i] > a[j] && longestIncLength[j]+1 > longestIncLength[i] {
				longestIncLength[i] = longestIncLength[j] + 1
			}
		}
	}

	longestDecLength := make([]int, cnt)
	longestDecLength[cnt-1] = 1
	for i := cnt - 2; i >= 0; i-- {
		longestDecLength[i] = 1
		for j := i + 1; j < cnt; j++ {
			if a[i] > a[j] && longestDecLength[j]+1 > longestDecLength[i] {
				longestDecLength[i] = longestDecLength[j] + 1
			}
		}
	}

	longestLen := 0
	for i := 0; i < cnt; i++ {
		l := 0
		if longestIncLength[i] > 1 && longestDecLength[i] > 1 {
			l = longestIncLength[i] + longestDecLength[i] - 1
		}
		if l > longestLen {
			longestLen = l
		}
	}

	return longestLen
}
