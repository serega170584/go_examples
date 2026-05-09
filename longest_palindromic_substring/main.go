package main

import "fmt"

func main() {
	s := []rune{'c', 'b', 'b', 'd'}
	if len(s) == 0 {
		fmt.Println([]rune{})
	}

	ms := []rune{s[0]}

	for i := 1; i < len(s); i++ {
		l := i - 1
		r := i + 1
		cs := []rune{s[i]}
		ms = getMaxPalindromicSubstring(s, l, r, cs, ms)
	}

	for i := 1; i < len(s); i++ {
		l := i - 1
		r := i
		var cs []rune
		ms = getMaxPalindromicSubstring(s, l, r, cs, ms)
	}

	for i := 0; i < len(s); i++ {
		l := i
		r := i + 1
		var cs []rune
		ms = getMaxPalindromicSubstring(s, l, r, cs, ms)
	}

	fmt.Println(string(ms))
}

func getMaxPalindromicSubstring(s []rune, l int, r int, cs []rune, ms []rune) []rune {
	func() {
		for l > 0 && r < len(s) {
			if s[l] != s[r] {
				return
			}

			cs = append([]rune{s[l]}, cs...)
			cs = append(cs, s[r])

			if len(cs) > len(ms) {
				ms = cs
			}

			l--
			r++
		}
	}()

	return ms
}
