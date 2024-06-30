package main

import (
	"fmt"
	"slices"
)

func main() {
	groups := []string{"sdad", "adsd", "ewq", "aaaa", "ab", "qwe", "wqe"}
	anagramGroups := getAnagramGroups(groups)
	fmt.Println(anagramGroups)
}

func getAnagramGroups(groups []string) [][]string {
	gm := make(map[string][]string, len(groups))
	for _, g := range groups {
		sl := []rune(g)
		slices.Sort(sl)
		gm[string(sl)] = append(gm[string(sl)], g)
	}

	ag := make([][]string, 0, len(groups))
	o := make([]string, 0, len(groups))
	for _, g := range gm {
		if len(g) > 1 {
			ag = append(ag, g)
		} else {
			o = append(o, g[0])
		}
	}

	if len(o) > 0 {
		ag = append(ag, o)
	}

	return ag
}
