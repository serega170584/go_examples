package main

import "fmt"

func main() {
	s := 100
	n := 10
	k := 0

	lv := 1
	rv := s

	leftMostIndex := k - 1
	rightMostIndex := k + 1

	lastIndex := n - 1

	for lv < rv {
		m := (lv+rv)/2 + 1

		mostValue := m - 1

		lp := 0
		if leftMostIndex >= 0 {
			lp = getLeftPart(mostValue, leftMostIndex)
		}

		rp := 0
		if rightMostIndex < n {
			rp = getRightPArt(mostValue, lastIndex, rightMostIndex)
		}

		t := lp + rp + m

		if t > s {
			rv = m - 1
			continue
		}

		lv = m

		fmt.Println(lv, rv)
	}

	fmt.Println(rv)
}

func getLeftPart(mostValue int, leftMostIndex int) int {
	lfv := 0
	if mostValue > leftMostIndex {
		lfv = mostValue - leftMostIndex
	}
	return (lfv + mostValue) * (mostValue - lfv + 1) / 2
}

func getRightPArt(mostValue int, lastIndex int, rightMostIndex int) int {
	rfv := 0
	indexDiff := lastIndex - rightMostIndex
	if mostValue > indexDiff {
		rfv = mostValue - indexDiff
	}
	return (rfv + mostValue) * (mostValue - rfv + 1) / 2
}
