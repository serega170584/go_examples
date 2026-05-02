package main

import (
	"fmt"
	"strings"
)

func main() {
	numRows := 3
	str := []rune("PAYPALISHIRING")

	rows := make([]strings.Builder, numRows)
	currRow := 1
	isDown := true
	isUp := false
	rows[0].WriteRune(str[0])
	for i := 1; i < len(str); i++ {
		rows[currRow].WriteRune(str[i])
		if currRow == numRows-1 || currRow == 0 {
			isDown = !isDown
			isUp = !isUp
		}

		if isDown {
			currRow++
		}

		if isUp {
			currRow--
		}
	}

	var res strings.Builder
	for j := 0; j < numRows; j++ {
		res.WriteString(rows[j].String())
	}

	fmt.Println(res.String())
}
