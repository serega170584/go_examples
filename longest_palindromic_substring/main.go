package main

import "fmt"

func main() {
	s := []rune{'b', 'a', 'b', 'a', 'd'}
	if len(s) == 0 {
		fmt.Println([]rune{})
	}

	res := []rune{s[0]}
	for i := 1; i < len(s); i++ {
		left := i - 1
		right := i + 1
		v := []rune{s[i]}
		for left >= 0 && right < len(s) {
			if s[left] != s[right] {
				break
			}

			v = append([]rune{s[left]}, v...)
			v = append(v, s[right])
			if len(v) > len(res) {
				res = v
			}

			left--
			right++
		}

		left = i - 1
		right = i
		v = []rune{}
		for left >= 0 && right < len(s) {
			if s[left] != s[right] {
				break
			}

			v = append([]rune{s[left]}, v...)
			v = append(v, s[right])
			if len(v) > len(res) {
				res = v
			}

			left--
			right++
		}
	}

	fmt.Println(string(res))
}
