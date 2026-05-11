package main

import "fmt"

func main() {
	s := []rune{'a', 'b', 'c', 'a', 'b', 'c', 'b', 'b'}

	im := make(map[rune]int)
	im[s[0]] = 0
	l := 0
	ml := 1
	cl := 1

	for r := 1; r < len(s); r++ {
		if v, ok := im[s[r]]; ok {
			if v >= l {
				l = v + 1
			}
		}
		cl = r - l + 1
		ml = max(cl, ml)

		im[s[r]] = r
	}

	fmt.Println(ml)
}
