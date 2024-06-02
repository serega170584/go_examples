package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(getAnagramGroups([]string{"abc", "cba", "adc", "dca", "cda", "a", "b", "c", "cd"}))
}

func getAnagramGroups(a []string) [][]string {
	groups := make(map[string][]string)
	for _, v := range a {
		key := []int32(v)
		slices.Sort(key)
		groups[string(key)] = append(groups[string(key)], v)
	}
	groupsSlice := make([][]string, 0)
	var newGroup []string
	for _, v := range groups {
		if len(v) == 1 {
			newGroup = append(newGroup, v[0])
		} else {
			groupsSlice = append(groupsSlice, v)
		}
	}
	if newGroup != nil {
		groupsSlice = append(groupsSlice, newGroup)
	}
	return groupsSlice
}
