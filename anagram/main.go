package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	s1 := scanner.Text()

	scanner.Scan()
	s2 := scanner.Text()

	cnt1 := make(map[rune]int, utf8.RuneCountInString(s1))
	for _, v := range s1 {
		cnt1[v]++
	}

	cnt2 := make(map[rune]int, utf8.RuneCountInString(s2))
	for _, v := range s2 {
		if _, ok := cnt1[v]; !ok {
			fmt.Println("NO")
			return
		}
		cnt2[v]++
	}

	for v := range cnt1 {
		if cnt1[v] != cnt2[v] {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
