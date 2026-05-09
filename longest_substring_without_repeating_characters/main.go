package main

import "fmt"

func main() {
	s := []rune{'b', 'b', 'b', 'b'}

	if len(s) == 0 {
		fmt.Println(0)
		return
	}

	if len(s) == 1 {
		fmt.Println(1)
		return
	}

	s = append([]rune{'0'}, s...)

	im := make(map[rune]int)

	l := 1
	lv := s[l]
	im[lv] = l

	r := 2

	ml := 1

	for r < len(s) {
		v := s[r]
		evi := im[v]

		if evi >= l {
			l = evi + 1
		}

		length := r - l + 1
		ml = max(ml, length)

		r++
		im[v] = r
	}

	fmt.Println(ml)
}
