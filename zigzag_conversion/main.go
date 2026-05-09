package main

import "fmt"

func main() {
	s := []rune{'P', 'A', 'Y', 'P', 'A', 'L', 'I', 'S', 'H', 'I', 'R', 'I', 'N', 'G'}
	numRows := 3

	sl := make([][]rune, numRows)
	down := true
	up := false
	sl[0] = []rune{s[0]}
	csi := 0

	for i := 1; i < len(s); i++ {
		if down {
			csi++
		}

		if up {
			csi--
		}

		sl[csi] = append(sl[csi], s[i])

		if csi == numRows-1 || csi == 0 {
			down = !down
			up = !up
		}
	}

	for _, v := range sl {
		fmt.Println(string(v))
	}
}
