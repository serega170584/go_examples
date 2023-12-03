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

	list := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Enter target:")
	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())

	var count, sum int
	for _, val := range list {
		sum += val

		if sum > target {
			break
		}

		count++
	}

	fmt.Println(count)
}
