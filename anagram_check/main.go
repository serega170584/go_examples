package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	first := scanner.Text()

	scanner.Scan()
	second := scanner.Text()

	firstSlice := []rune(first)
	secondSlice := []rune(second)

	if len(firstSlice) != len(secondSlice) {
		fmt.Println("NO")
	}

	firstMap := make(map[rune]int, len(firstSlice))
	for _, v := range firstSlice {
		firstMap[v]++
	}

	secondMap := make(map[rune]int, len(secondSlice))
	for _, v := range secondSlice {
		secondMap[v]++
	}

	for i, v := range firstMap {
		if v != secondMap[i] {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")

}
