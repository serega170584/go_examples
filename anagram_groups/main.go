package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(getAnagramGroups([]string{"abc", "cba", "adc", "dca", "cda", "a", "b", "c", "cd"}))
}

func getAnagramGroups(a []string) [][]string {
	am := make(map[string][]string, len(a))
	for _, v := range a {
		r := []int32(v)
		slices.Sort(r)
		rv := string(r)
		am[rv] = append(am[rv], v)
	}

	only := make([]string, 0)
	groups := make([][]string, 0)
	for _, v := range am {
		if len(v) == 1 {
			only = append(only, v[0])
		} else {
			groups = append(groups, v)
		}
	}

	if len(only) != 0 {
		groups = append(groups, only)
	}

	return groups
}
