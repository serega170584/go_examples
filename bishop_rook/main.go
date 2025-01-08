package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	list := make([]string, 0, 8)
	for i := 0; i < 8; i++ {
		scanner.Scan()
		list = append(list, scanner.Text())
	}

	bottom := make([][]bool, 0, 8)
	right := make([][]bool, 0, 8)
	crossLeft := make([][]bool, 0, 8)
	crossRight := make([][]bool, 0, 8)
	for i, v := range list {
		bottom = append(bottom, make([]bool, 0, 8))
		right = append(right, make([]bool, 0, 8))
		crossLeft = append(crossLeft, make([]bool, 0, 8))
		crossRight = append(crossRight, make([]bool, 0, 8))
		j := 0
		for _, s := range v {
			crossLeftV := false
			if i > 0 && j > 0 && crossLeft[i-1][j-1] {
				crossLeftV = true
			}
			if s == 'R' {
				crossLeftV = false
			}
			if s == 'B' {
				crossLeftV = true
			}
			crossLeft[i] = append(crossLeft[i], crossLeftV)

			crossRightV := false
			if i > 0 && j < 7 && crossRight[i-1][j+1] {
				crossRightV = true
			}
			if s == 'R' {
				crossRightV = false
			}
			if s == 'B' {
				crossRightV = true
			}
			crossRight[i] = append(crossRight[i], crossRightV)

			rightV := false
			if j > 0 && right[i][j-1] {
				rightV = true
			}
			if s == 'B' {
				rightV = false
			}
			if s == 'R' {
				rightV = true
			}
			right[i] = append(right[i], rightV)

			bottomV := false
			if i > 0 && bottom[i-1][j] {
				bottomV = true
			}
			if s == 'B' {
				bottomV = false
			}
			if s == 'R' {
				bottomV = true
			}
			bottom[i] = append(bottom[i], bottomV)

			j++
		}
	}

	top := make([][]bool, 0, 8)
	for i := 0; i < 8; i++ {
		top = append(top, make([]bool, 8))
	}

	left := make([][]bool, 0, 8)
	for i := 0; i < 8; i++ {
		left = append(left, make([]bool, 8))
	}

	crossLeftTop := make([][]bool, 0, 8)
	for i := 0; i < 8; i++ {
		crossLeftTop = append(crossLeftTop, make([]bool, 8))
	}

	crossRightTop := make([][]bool, 0, 8)
	for i := 0; i < 8; i++ {
		crossRightTop = append(crossRightTop, make([]bool, 8))
	}

	for i := 7; i >= 0; i-- {
		j := 7
		v := list[i]
		reverse := make([]rune, 8)
		for ind, s := range v {
			reverse[7-ind] = s
		}

		for _, s := range reverse {
			crossLeftTopV := false
			if i < 7 && j > 0 && crossLeftTop[i+1][j-1] {
				crossLeftTopV = true
			}
			if s == 'R' {
				crossLeftTopV = false
			}
			if s == 'B' {
				crossLeftTopV = true
			}
			crossLeftTop[i][j] = crossLeftTopV

			crossRightTopV := false
			if i < 7 && j < 7 && crossRightTop[i+1][j+1] {
				crossRightTopV = true
			}
			if s == 'R' {
				crossRightTopV = false
			}
			if s == 'B' {
				crossRightTopV = true
			}
			crossRightTop[i][j] = crossRightTopV

			leftV := false
			if j < 7 && left[i][j+1] {
				leftV = true
			}
			if s == 'B' {
				leftV = false
			}
			if s == 'R' {
				leftV = true
			}
			left[i][j] = leftV

			topV := false
			if i < 7 && top[i+1][j] {
				topV = true
			}
			if s == 'B' {
				topV = false
			}
			if s == 'R' {
				topV = true
			}
			top[i][j] = topV

			j--
		}
	}

	cnt := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if !bottom[i][j] && !right[i][j] && !crossLeft[i][j] && !crossRight[i][j] && !crossLeftTop[i][j] && !crossRightTop[i][j] && !top[i][j] && !left[i][j] {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}
