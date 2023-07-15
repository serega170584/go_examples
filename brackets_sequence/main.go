package main

import (
	"fmt"
	"log"
)

func main() {
	var cnt int
	_, err := fmt.Scan(&cnt)
	if err != nil {
		log.Fatal(err)
	}

	if cnt == 1 {
		fmt.Println("{}")
		return
	}

	prev := []string{"{}"}
	lastPrevIndex := 0
	cur := make([]string, 0)
	for i := 1; i < cnt; i++ {
		curIndex := 0
		for prevIndex, val := range prev {
			if curIndex == len(cur) {
				cur = append(cur, "")
			}
			cur[curIndex] = "{" + val + "}"
			curIndex++
			if curIndex == len(cur) {
				cur = append(cur, "")
			}
			cur[curIndex] = val + "{}"
			curIndex++
			if prevIndex != lastPrevIndex {
				if curIndex == len(cur) {
					cur = append(cur, "")
				}
				cur[curIndex] = "{}" + val
				curIndex++
			}
		}

		for curI, val := range cur {
			if curI == len(prev) {
				prev = append(prev, val)
			}
			prev[curI] = val
		}
		lastPrevIndex = len(prev) - 1
	}
	fmt.Println(cur)
}
