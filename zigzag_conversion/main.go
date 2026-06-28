package main

import "fmt"

func main() {
	s := []rune{'P', 'A', 'Y', 'P', 'A', 'L', 'I', 'S', 'H', 'I', 'R', 'I', 'N', 'G'}
	numRows := 3

	if numRows == 1 {
		fmt.Println(s)
	}

	moveUp := false
	moveDown := true
	sl := make([][]rune, numRows)
	ri := 0
	for i := 0; i < len(s); i++ {
		if ri == -1 {
			ri = 1
			moveUp, moveDown = moveDown, moveUp
		}

		if ri == numRows {
			ri = numRows - 2
			moveUp, moveDown = moveDown, moveUp
		}

		sl[ri] = append(sl[ri], s[i])

		if moveDown {
			ri++
		}

		if moveUp {
			ri--
		}
	}

	res := string(sl[0])
	for i := 1; i < numRows; i++ {
		res += string(sl[i])
	}

	fmt.Println(res)
}
