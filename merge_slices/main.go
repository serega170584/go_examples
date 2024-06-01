package main

import "fmt"

func main() {
	fmt.Println(mergeSlices([]int{1, 2, 6}, []int{3, 3, 4, 5}))
}

func mergeSlices(a []int, b []int) []int {
	al := len(a)
	bl := len(b)
	var mainS, foreignS []int
	foreignL := 0
	mainL := 0
	if al > bl {
		mainS = a
		foreignS = b
		foreignL = bl
		mainL = al
	} else {
		mainS = b
		foreignS = a
		foreignL = al
		mainL = bl
	}

	for i := 0; i < foreignL; i++ {
		mainS = append(mainS, 0)
	}

	mainP := mainL - 1
	foreignP := foreignL - 1
	for i := mainL + foreignL - 1; i >= 0; i-- {
		if mainP == -1 {
			mainS[i] = foreignS[foreignP]
			foreignP--
			continue
		}

		if foreignP == -1 {
			mainS[i] = mainS[mainP]
			mainP--
			continue
		}

		if mainS[mainP] > foreignS[foreignP] {
			mainS[i] = mainS[mainP]
			mainP--
		} else {
			mainS[i] = foreignS[foreignP]
			foreignP--
		}
	}

	return mainS
}
