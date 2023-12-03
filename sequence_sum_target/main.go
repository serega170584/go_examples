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

	fmt.Println("Enter target ")
	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())

	var right, sum int
	for left := range list {

		if left > 0 {
			sum -= list[left-1]
		}

		if sum == target {
			fmt.Println("Yes")
			return
		}

		for right < len(list) && sum < target {
			sum += list[right]
			right++
		}
	}

	fmt.Println("No")
	return
}
