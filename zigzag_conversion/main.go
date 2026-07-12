package main

import "fmt"

func main() {
	s := []rune{'P', 'A', 'Y', 'P', 'A', 'L', 'I', 'S', 'H', 'I', 'R', 'I', 'N', 'G'}
	numRows := 3

	if numRows == 1 {
		fmt.Println(s)
	}

	down := true
	up := false

	j := 0

	sl := make([][]rune, numRows)
	for i := 0; i < len(s); i++ {
		sl[j] = append(sl[j], s[i])

		if j == 0 {
			down = true
			up = false
		}

		if j == numRows-1 {
			down = false
			up = true
		}

		if up {
			j--
		}

		if down {
			j++
		}
	}

	res := ""
	for i := 0; i < len(sl); i++ {
		res = res + string(sl[i])
	}

	fmt.Println(res)
}
