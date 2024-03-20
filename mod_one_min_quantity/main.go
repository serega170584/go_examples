package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
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

	fmt.Println(getModOneMinCnt(n, list))
}

func getModOneMinCnt(n int, list []int) int {
	numCnts := make(map[int]int, n)
	keys := make([]int, 0, n)
	cnt := 0
	for _, v := range list {
		if _, ok := numCnts[v]; !ok {
			keys = append(keys, v)
		}
		numCnts[v]++
		cnt++
	}

	keysLen := len(keys)
	if keysLen == 1 {
		return 0
	}

	slices.Sort(keys)
	minExcludeCnt := math.MaxInt
	prev := keys[0]
	for i := 1; i < keysLen; i++ {
		v := keys[i]
		if v-prev == -1 || v-prev == 1 {
			minExcludeCnt = min(cnt-numCnts[v]-numCnts[prev], minExcludeCnt)
		} else {
			minExcludeCnt = min(cnt-numCnts[v], minExcludeCnt)
		}
		prev = keys[i]
	}

	return minExcludeCnt
}
