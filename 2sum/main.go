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

	dict := make(map[int]int, cnt)
	for i, val := range list {
		dict[val] = i
		addVal := target - val

		if existed, ok := dict[addVal]; ok {
			fmt.Println(i, " ", existed)
		}
	}
}
