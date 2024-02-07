package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	countScanner := bufio.NewScanner(os.Stdin)
	countScanner.Split(bufio.ScanWords)

	bytesScanner := bufio.NewScanner(os.Stdin)
	bytesScanner.Split(bufio.ScanBytes)

	fmt.Println("Enter dictionary words count")
	countScanner.Scan()
	dictWordsCnt, _ := strconv.Atoi(countScanner.Text())

	dict := make([][]byte, dictWordsCnt)
	for i := 0; i < dictWordsCnt; i++ {
		fmt.Println("Enter word length")
		countScanner.Scan()
		wordLength, _ := strconv.Atoi(countScanner.Text())
		dict[i] = make([]byte, wordLength)
		for j := 0; j < wordLength; j++ {
			bytesScanner.Scan()
			dict[i][j] = bytesScanner.Bytes()[0]
		}
	}

	fmt.Println("Enter text count")
	countScanner.Scan()
	textCnt, _ := strconv.Atoi(countScanner.Text())

	fmt.Println("Enter text")
	text := make([]string, textCnt)
	for i := 0; i < textCnt; i++ {
		countScanner.Scan()
		text[i] = countScanner.Text()
	}

	fmt.Println("Searched words:", wordsInDict(dictWordsCnt, textCnt, dict, text))
}

func wordsInDict(wordsCount int, textCount int, dict [][]byte, text []string) []string {
	words := make(map[string]struct{}, wordsCount)
	for _, word := range dict {
		wordBytesCnt := len(word)
		for i := 0; i < wordBytesCnt; i++ {
			var sb strings.Builder

			leftWords := make([]byte, i)
			copy(leftWords, word[0:i])
			sb.Write(leftWords)

			rightWords := make([]byte, wordBytesCnt-i-1)
			copy(rightWords, word[i+1:wordBytesCnt])
			sb.Write(rightWords)

			words[sb.String()] = struct{}{}

		}
	}

	searched := make([]string, textCount)
	ind := 0
	for _, val := range text {
		if _, ok := words[val]; ok {
			searched[ind] = val
			ind++
		}
	}
	return searched[0:ind]
}
