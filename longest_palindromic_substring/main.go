package main

import "fmt"

func main() {
	s := []rune{'c', 'b', 'b', 'd'}
	res := []rune{s[0]}

	for i := 1; i < len(s); i++ {
		v := []rune{s[i]}
		left := i - 1
		right := i + 1
		for left >= 0 && right < len(s) {
			if s[left] != s[right] {
				break
			}

			v = append([]rune{s[left]}, v...)
			v = append(v, []rune{s[right]}...)

			left--
			right++
		}

		if len(v) > len(res) {
			res = v
		}

		left = i - 1
		right = i
		v = []rune{}
		for left >= 0 && right < len(s) {
			if s[left] != s[right] {
				break
			}

			v = append([]rune{s[left]}, v...)
			v = append(v, []rune{s[right]}...)

			left--
			right++
		}

		if len(v) > len(res) {
			res = v
		}
	}

	fmt.Println(string(res))
}
