package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	str []rune
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)

	intScanner := bufio.NewScanner(os.Stdin)
	intScanner.Split(bufio.ScanWords)

	fmt.Println("Enter string length")
	intScanner.Scan()
	cnt, _ := strconv.Atoi(intScanner.Text())

	fmt.Println("Enter string")
	word := make([]rune, cnt)

	for i := 0; i < cnt; i++ {
		scanner.Scan()
		word[i] = []rune(scanner.Text())[0]
	}

	minStr, minLength := minWordPackingSequence(cnt, word)

	fmt.Println("Got packing string:", minStr)
	fmt.Println("Got min length", minLength)
}

func minWordPackingSequence(cnt int, word []rune) (string, int) {
	dp := make([][]*Node, cnt)
	minLength := make([][]int, cnt)
	for i := 0; i < cnt; i++ {
		dp[i] = make([]*Node, cnt)
		str := make([]rune, 1)
		str[0] = word[i]
		dp[i][i] = &Node{str: str}
		minLength[i] = make([]int, cnt)
		minLength[i][i] = 1
	}

	for j := 1; j < cnt; j++ {
		for i := 0; i < cnt-j; i++ {
			left := i
			right := i + j
			length := j + 1
			str := make([]rune, length)
			copy(str, word[left:right+1])

			minStrLength := length
			minStr := make([]rune, length)
			copy(minStr, str)

			for k := 0; k < j; k++ {
				if length%(k+1) == 0 {
					partsCnt := length / (k + 1)

					partStr, partLength := getPartMinStr(partsCnt, k+1, length, str)

					if partLength < minStrLength {
						minStrLength = partLength
						copy(minStr, partStr)
					}
				}

				strLength := minLength[left][left+k] + minLength[left+k+1][right]
				if strLength < minStrLength {
					minStrLength = strLength
					curStr := make([]rune, minLength[left][left+k])
					curStr = dp[left][left+k].str
					curStr = append(curStr, dp[left+k+1][right].str[0:minLength[left+k+1][right]]...)
					copy(minStr, curStr)
				}
			}

			dp[left][right] = &Node{str: minStr[0:minStrLength]}
			minLength[left][right] = minStrLength
		}
	}

	return string(dp[0][cnt-1].str[0:minLength[0][cnt-1]]), minLength[0][cnt-1]
}

func getPartMinStr(partsCnt int, partLength int, length int, str []rune) ([]rune, int) {
	part := make([]rune, partLength)
	for i := 0; i < partLength; i++ {
		part[i] = str[i]
		for j := 1; j < partsCnt; j++ {
			if str[i+j*partLength] != str[i+(j-1)*partLength] {
				return str, length
			}
		}
	}

	resLength := 0

	prefix := []rune(strconv.Itoa(partsCnt))
	prefixLength := len(prefix)

	resLength += prefixLength + 2 + partLength

	res := make([]rune, resLength)

	for i := 0; i < prefixLength; i++ {
		res[i] = prefix[i]
	}

	res[prefixLength] = []rune("(")[0]

	for i := 0; i < partLength; i++ {
		res[prefixLength+1+i] = part[i]
	}

	res[resLength-1] = []rune(")")[0]

	return res, resLength
}
