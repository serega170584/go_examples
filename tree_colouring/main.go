package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// lp1_____lp2
//   rp1___rp2
//     rp1_____rp2

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	p1, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	i1, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	p2, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	i2, _ := strconv.Atoi(scanner.Text())

	fmt.Println(getColouredTreesCnt(p1, i1, p2, i2))

}

func getColouredTreesCnt(p1 int, i1 int, p2 int, i2 int) int {
	cp1 := min(p1-i1, p1+i1)
	cp2 := max(p1-i1, p1+i1)

	cp3 := min(p2-i2, p2+i2)
	cp4 := max(p2-i2, p2+i2)

	cnt := cp2 - cp1 + 1
	cnt += cp4 - cp3 + 1

	if cp3 > cp1 {
		return getTreeCnt(cp2, cp3, cp4, cnt)
	}

	return getTreeCnt(cp4, cp1, cp2, cnt)
}

func getTreeCnt(lp2 int, rp1 int, rp2 int, cnt int) int {
	n := 0
	if rp1 > lp2 {
		return cnt
	} else {
		n = cnt - lp2 + rp1 - 1
	}

	if lp2 > rp2 {
		n += lp2 - rp2
	}

	return n
}
