package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)

	intScanner := bufio.NewScanner(os.Stdin)
	intScanner.Split(bufio.ScanWords)

	fmt.Println("Enter sequence count")
	intScanner.Scan()
	cnt, _ := strconv.Atoi(intScanner.Text())

	fmt.Println("Enter sequence")
	sequence := make([]rune, cnt)
	scanner.Scan()
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		sequence[i] = []rune(scanner.Text())[0]
	}

	fmt.Println("Got min brackets count", minBracketSequenceLength(cnt, sequence))
}

func minBracketSequenceLength(cnt int, str []rune) int {
	dp := make([][]int, cnt)
	for i := 0; i < cnt; i++ {
		dp[i] = make([]int, cnt)
		for j := 0; j < cnt; j++ {
			dp[i][j] = math.MaxInt
			if i == j {
				dp[i][j] = 1
			}
		}
	}

	bracketAssociations := make(map[string]string, 3)
	bracketAssociations["{"] = "}"
	bracketAssociations["("] = ")"
	bracketAssociations["["] = "]"

	for i := 1; i < cnt; i++ {
		for j := 0; j < cnt-i; j++ {
			for k := 0; k < i; k++ {
				left := j
				right := j + i

				dp[j][j+i] = min(dp[j][j+i], dp[j][j+k]+dp[j+k+1][j+i])

				innerLeft := left + 1
				innerRight := right - 1
				innerLength := innerRight - innerLeft
				innerCost := 0
				if innerLength >= 0 {
					innerCost = dp[innerLeft][innerRight]
				}

				leftSym := string(str[left])
				rightSym := string(str[right])
				if bracketAssociations[leftSym] == rightSym {
					dp[j][j+i] = min(dp[j][j+i], innerCost)
				}
			}
		}
	}

	return dp[0][cnt-1]
}
