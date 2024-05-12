package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(getAnagramGroups([]string{"abc", "cba", "adc", "dca", "cda", "a", "b", "c", "cd"}))
}

func getAnagramGroups(a []string) [][]string {
	l := len(a)
	groups := make([][]string, 0, l)
	gm := make(map[string][]string, l)
	for _, v := range a {
		var sb strings.Builder
		r := []int32(v)
		slices.Sort(r)

		for _, rv := range r {
			sb.WriteString(string(rv))
		}

		k := sb.String()
		if _, ok := gm[k]; !ok {
			gm[k] = make([]string, 0, l)
		}
		gm[k] = append(gm[k], v)
	}

	eg := make([]string, 0, l)
	for _, v := range gm {
		if len(v) == 1 {
			eg = append(eg, v[0])
		} else {
			groups = append(groups, []string{})
			gl := len(groups)
			groups[gl-1] = make([]string, len(v))
			groups[gl-1] = v
		}
	}

	groups = append(groups, []string{})
	gl := len(groups)
	groups[gl-1] = make([]string, len(eg))
	groups[gl-1] = eg

	return groups
}
