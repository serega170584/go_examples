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

	a := make([][]int, cnt)
	for i := 0; i < cnt; i++ {
		a[i] = make([]int, cnt)
		for j := 0; j < cnt; j++ {
			scanner.Scan()
			a[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}

	rotate(a)

	for i := 0; i < cnt; i++ {
		fmt.Println(a[i])
	}
}

func rotate(a [][]int) {
	l := len(a)
	if l == 0 || l == 1 {
		return
	}

	for layer := 0; layer < l/2; layer++ {
		first := layer
		last := l - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			top := a[first][i]

			a[first][i] = a[last-offset][first]

			a[last-offset][first] = a[last][last-offset]

			a[last][last-offset] = a[i][last]

			a[i][last] = top
		}
	}
}
