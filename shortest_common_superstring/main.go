package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// "bwpiqhiym", "yalcyea", "vxggfitknygyv", "xnspubqjppjbrl", "ugesmmxwjjlk", "mgbnwvf", "yveolprfdcajiu"

// 1 2 3 4 5 6 7
// 12 21 13 31 14 41 15 51 16 61 17 71
// 123 312 213 321 132 231
// 4123 1234 4312 3124 4213 2134 4321 3214

// bwpiqhiymyalcyea yalcyeabwpiqhiym
// bwpiqhiymyalcyeavxggfitknygyv vxggfitknygyvbwpiqhiymyalcyea yalcyeabwpiqhiymvxggfitknygyv vxggfitknygyvalcyeabwpiqhiym
// bwpiqhiymyalcyeavxggfitknygyvxnspubqjppjbrl xnspubqjppjbrlbwpiqhiymyalcyeavxggfitknygyv

//   a b c d e f g
// e 0 0 0 0 1 0 0
// f 0 0 0 0 0 2 0
// g 0 0 0 0 0 0 3

//   a b c d e f g
// d 0 0 0 1 0 0 0
// e 0 0 0 0 2 0 0
// f 0 0 0 0 0 3 0

//   a b c d e f g
// e 0 0 0 0 1 0 0
// f 0 0 0 0 0 2 0
// g 0 0 0 0 0 0 3
// h 0 0 0 0 0 0 0

//   a b c d e f g i
// e 0 0 0 0 1 0 0 0
// f 0 0 0 0 0 2 0 0
// g 0 0 0 0 0 0 3 0
// h 0 0 0 0 0 0 0 0

//   a b c d e f g
// e 0 0 0 0 1 0 0
// f 0 0 0 0 0 2 0
// g 0 0 0 0 0 0 3

//   a b c d e f g
// x 0 0 0 0 0 0 0
// a 1 0 0 0 0 0 0
// b 0 2 0 0 0 0 0

//   a b a b e f g
// x 0 0 0 0 0 0 0
// a 1 0 1 0 0 0 0
// b 0 2 0 2 0 0 0

//   a b c d e f g e f g
// e 0 0 0 0 1 0 0 1 0 0
// f 0 0 0 0 0 2 0 0 2 0
// g 0 0 0 0 0 0 3 0 0 3

//   a b c d e f g e f g
// e 0 0 0 0 1 0 0 1 0 0
// f 0 0 0 0 0 2 0 0 2 0
// g 0 0 0 0 0 0 3 0 0 3
// h 0 0 0 0 0 0 0 0 0 0

// 1 2 3 [[1 2 3]] 1 2 3
// 1 3 2 [[1 2 3] [1 3 2]] 1 2 3
// 2 1 3 [[1 2 3] [1 3 2] [2 1 3]]
// 2 3 1 [[1 2 3] [1 3 2] [2 1 3] [2 3 1]] 2 1 3
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	list := make([]string, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		list[i] = scanner.Text()
	}

	indexPermutations := getIndexPermutations(list)

	curList := make([]string, len(list))
	for i, v := range indexPermutations[0] {
		curList[i] = list[v]
	}
	ml := getListSuperStringMInLength(curList)

	for i := 1; i < len(indexPermutations); i++ {
		curList = make([]string, len(list))
		ip := indexPermutations[i]
		for j, v := range ip {
			curList[j] = list[v]
		}

		curMl := getListSuperStringMInLength(curList)
		if curMl < ml {
			ml = curMl
		}
	}

	fmt.Println(ml)
}

func getIndexPermutations(list []string) [][]int {
	permutations := [][]int{{0}}
	for i := 1; i < len(list); i++ {
		curPermutations := make([][]int, 0)
		for _, perm := range permutations {
			curPerm := append(perm, i)
			curPermutations = append(curPermutations, curPerm)
			for j := range curPerm {
				var addPerm = make([]int, len(curPerm))
				copy(addPerm, curPerm)
				addPerm[j], addPerm[len(curPerm)-1] = addPerm[len(curPerm)-1], addPerm[j]
				curPermutations = append(curPermutations, addPerm)
			}
		}
		permutations = curPermutations
	}

	return permutations
}

func getListSuperStringMInLength(list []string) int {
	runeList := make([][]rune, 0)
	for _, v := range list {
		runeList = append(runeList, []rune(v))
	}

	return getRunesSuperStringMinLength(runeList)
}

func getRunesSuperStringMinLength(runeList [][]rune) int {
	superString := runeList[0]

	for i := 1; i < len(runeList); i++ {
		superStringLen := len(superString)

		str := runeList[i]
		strLen := len(str)

		dp := make([][]int, superStringLen+1)
		for k := 0; k <= superStringLen; k++ {
			dp[k] = make([]int, strLen+1)
		}

		foundCnt := 0

		for sui := 1; sui < superStringLen; sui++ {
			for si := 1; si <= strLen; si++ {
				if superString[sui-1] == str[si-1] {
					dp[sui][si] = dp[sui-1][si-1] + 1
				}
				if dp[sui][si] == strLen {
					foundCnt = strLen
					break
				}
			}

			if foundCnt == strLen {
				break
			}
		}

		if foundCnt == strLen {
			continue
		}

		for si := 1; si <= strLen; si++ {
			if superString[superStringLen-1] == str[si-1] {
				dp[superStringLen][si] = dp[superStringLen-1][si-1] + 1
				foundCnt = dp[superStringLen][si]
			}
		}

		superString = append(superString[0:superStringLen-foundCnt], str...)
	}

	return len(superString)
}
