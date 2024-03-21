package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := makeScanner()

	scanner.Scan()
	s := strings.Split(scanner.Text(), " ")
	dn := len(s)
	dict := make([][]rune, dn)
	for i, v := range s {
		dict[i] = []rune(v)
	}

	scanner.Scan()
	ws := strings.Split(scanner.Text(), " ")
	wn := len(ws)
	words := make([][]rune, wn)
	for i, v := range ws {
		words[i] = []rune(v)
	}

	rs := make([]string, wn)
	r := getReplacedWords(dn, dict, wn, words)
	for i, v := range r {
		rs[i] = string(v)
	}

	fmt.Println(strings.Join(rs, " "))
}

func getReplacedWords(dn int, dict [][]rune, wn int, words [][]rune) [][]rune {
	dictMap := make(map[string]struct{}, dn)
	for _, v := range dict {
		dictMap[string(v)] = struct{}{}
	}

	replaced := make([][]rune, 0, wn)
	for _, word := range words {
		prefix := make([]rune, 0, len(word))
		for _, v := range word {
			prefix = append(prefix, v)
			if _, ok := dictMap[string(prefix)]; ok {
				break
			}
		}
		replaced = append(replaced, prefix)
	}

	return replaced
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}
