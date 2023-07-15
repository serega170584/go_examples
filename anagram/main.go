package main

import "fmt"

func main() {
	a := "aaabbaaaccdddbbfffbbaaadd"
	b := "aababbaaacdddbfffcbbaaadad"
	cnt := len(a)
	aCounts := make(map[int32]int, cnt)
	for _, val := range a {
		if _, ok := aCounts[val]; !ok {
			aCounts[val] = 0
		}
		aCounts[val]++
	}

	bCounts := make(map[int32]int, cnt)
	for _, val := range b {
		if _, ok := bCounts[val]; !ok {
			bCounts[val] = 0
		}
		bCounts[val]++
	}

	for key, val := range aCounts {
		if val != bCounts[key] {
			fmt.Println("a != b")
			return
		}
	}
	fmt.Println("a = b")
}
