package main

import (
	"fmt"
	"strconv"
	"strings"
)

// RDRU
// [] R [] D [] R [] U []
// [0 0] [1 0] [1 1] [2 1] [2 0]
func main() {
	a := []string{"U", "L", "D", "R", "U", "L", "D", "R"}
	fmt.Println(getLongestCycle(a))
}

func getLongestCycle(a []string) int {
	al := len(a)
	m := make(map[string]int, al)

	left := 0
	top := 0
	m[strings.Join([]string{strconv.Itoa(left), strconv.Itoa(top)}, "_")] = 0
	mc := 0
	for i, v := range a {
		if v == "L" {
			left--
		}

		if v == "R" {
			left++
		}

		if v == "U" {
			top--
		}

		if v == "D" {
			top++
		}

		k := strings.Join([]string{strconv.Itoa(left), strconv.Itoa(top)}, "_")

		if _, ok := m[k]; ok {
			mc = max(mc, i-m[k]+1)
		} else {
			m[k] = i + 1
		}
	}

	return mc
}
