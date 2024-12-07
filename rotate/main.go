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

	arr := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, make([]int, 0, n))
		for j := 0; j < n; j++ {
			scanner.Scan()
			v, _ := strconv.Atoi(scanner.Text())
			arr[i] = append(arr[i], v)
		}
	}

	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			top := arr[first][i]

			arr[first][i] = arr[last-offset][first]
			arr[last-offset][first] = arr[last][last-offset]
			arr[last][last-offset] = arr[i][last]
			arr[i][last] = top
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(arr[i])
	}
}
