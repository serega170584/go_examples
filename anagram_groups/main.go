package main

import (
	"fmt"
	"sort"
)

func main() {
	strings := []string{"sdad", "adsd", "ewq", "aaaa", "ab", "qwe", "wqe"}

	m := make(map[string][]string)
	for i := 0; i < len(strings); i++ {
		str := strings[i]
		sr := []rune(str)
		sort.Slice(sr, func(i int, j int) bool {
			return sr[i] < sr[j]
		})
		m[string(sr)] = append(m[string(sr)], str)
	}

	groups := make([][]string, 0, len(strings))
	oneGroup := make([]string, 0, len(strings))

	for _, v := range m {
		if len(v) > 1 {
			groups = append(groups, v)
			continue
		}

		oneGroup = append(oneGroup, v[0])
	}

	groups = append(groups, oneGroup)

	fmt.Println(groups)
}
