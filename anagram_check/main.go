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
	a := []rune(scanner.Text())

	scanner.Scan()
	b := []rune(scanner.Text())

	if len(a) != len(b) {
		fmt.Println("NO")
	}

	aMap := make(map[rune]int, len(a))
	bMap := make(map[rune]int, len(b))

	for i := 0; i < len(a); i++ {
		aMap[a[i]]++
		bMap[b[i]]++
	}

	for _, v := range a {
		if _, ok := bMap[v]; !ok {
			fmt.Println("NO")
			return
		}

		if aMap[v] != bMap[v] {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
