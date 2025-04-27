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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	s := 0
	for i := 0; i < n; i++ {
		s += a[i]
	}

	r := make([]int, s+1)
	for _, v := range a {
		r[v]++
	}

	ai := 0
	for i, v := range r {
		for j := 0; j < v; j++ {
			a[ai] = i
			ai++
		}
	}

	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, s+1)
	}
	dp[0][s] = []int{}

	for i := n; i >= 1; i-- {
		for j := n - i + 1; j >= 1; j-- {
			for k := 0; k <= s-a[i-1]; k++ {
				if dp[j-1][k+a[i-1]] != nil {
					dp[j][k] = append([]int{a[i-1]}, dp[j-1][k+a[i-1]]...)
				}
			}
		}
	}

	fp := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		if (i*s)%n == 0 && dp[i][s-i*s/n] != nil {
			fp = dp[i][s-i*s/n]
			break
		}
	}

	sp := make([]int, 0, n)

	ai = 0
	fpi := 0
	for fpi != len(fp) {
		if a[ai] == fp[fpi] {
			fpi++
		} else {
			sp = append(sp, a[ai])
		}
		ai++
	}

	for i := ai; i < n; i++ {
		sp = append(sp, a[i])
	}

	if len(sp) == 0 {
		fmt.Println(nil)
	}

	fmt.Println([][]int{fp, sp})

}
