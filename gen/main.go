package main

import "fmt"

func main() {
	generate("", 0, 0, 4)
}

func generate(s string, open int, close int, n int) {
	if len(s) == 2*n {
		fmt.Println(s)
		return
	}

	if open < n {
		generate(s+"(", open+1, close, n)
	}

	if open > close {
		generate(s+")", open, close+1, n)
	}
}
