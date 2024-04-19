package main

import "fmt"

func main() {
	generate("", 0, 0, 6)
}

func generate(cur string, open int, closed int, n int) {
	//fmt.Println(len(cur))
	if len(cur) == 2*n {
		fmt.Println(cur)
	}

	if open < n {
		generate(cur+"(", open+1, closed, n)
	}

	if closed < open {
		generate(cur+")", open, closed+1, n)
	}

}
