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

	fmt.Println("Enter first word length")
	scanner.Scan()
	firstWordLength, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter first word")
	firstWord := make([]rune, firstWordLength)
	scanner.Scan()
	for i := 0; i < firstWordLength; i++ {
		scanner.Scan()
		firstWord[i] = []rune(scanner.Text())[0]
	}

	scanner.Scan()

	fmt.Println("Enter second word length")
	scanner.Scan()
	secondWordLength, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter second word")
	secondWord := make([]rune, secondWordLength)
	scanner.Scan()
	for i := 0; i < secondWordLength; i++ {
		scanner.Scan()
		secondWord[i] = []rune(scanner.Text())[0]
	}

	fmt.Println("Got levenstein distance", levensteinDistance(firstWordLength, firstWord, secondWordLength, secondWord))
}

func levensteinDistance(firstWordLength int, firstWord []rune, secondWordLength int, secondWord []rune) int {
	dp := make([][]int, firstWordLength)
	for i := 0; i < firstWordLength; i++ {
		dp[i] = make([]int, secondWordLength)
		for j := 0; j < secondWordLength; j++ {
			dp[i][j] = math.MaxInt
		}
	}
	dp[0][0] = 0
	if firstWord[0] != secondWord[0] {
		dp[0][0] = 1
	}
	for i := range firstWord {
		for j := range secondWord {
			prevI := i - 1
			prevJ := j - 1

			addNum := 0
			if firstWord[i] != secondWord[j] {
				addNum = 1
			}

			repl := math.MaxInt
			if prevI >= 0 && prevJ >= 0 {
				repl = dp[prevI][prevJ] + addNum
			}

			ins := math.MaxInt
			if prevI >= 0 {
				ins = dp[prevI][j] + addNum
			}

			del := math.MaxInt
			if prevJ >= 0 {
				del = dp[i][prevJ] + addNum
			}
			dp[i][j] = min(dp[i][j], repl, ins, del)
		}
	}
	return dp[firstWordLength-1][secondWordLength-1]
}
