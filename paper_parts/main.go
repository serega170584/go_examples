package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//file, _ := os.Open("paper_parts/file.txt")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	length, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	firstWords := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		firstWords[i], _ = strconv.Atoi(scanner.Text())
	}

	secondWords := make([]int, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		secondWords[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(getMinPaperHeight(length, n, firstWords, m, secondWords))

}

func getMinPaperHeight(length int, n int, firstWords []int, m int, secondWords []int) int {
	maxFirstWordsLength := 0
	for _, v := range firstWords {
		maxFirstWordsLength = max(maxFirstWordsLength, v)
	}

	firstWordsHeight := 1
	curLength := firstWords[0]
	for i, v := range firstWords {
		if i == 0 {
			continue
		}
		newLength := curLength + v + 1
		if newLength > maxFirstWordsLength {
			firstWordsHeight++
			curLength = v
		} else {
			curLength = newLength
		}
	}

	maxSecondWordsLength := length - maxFirstWordsLength

	secondWordsHeight := 1
	curLength = secondWords[0]
	for i, v := range secondWords {
		if i == 0 {
			continue
		}
		newLength := curLength + v + 1
		if newLength > maxSecondWordsLength {
			secondWordsHeight++
			curLength = v
		} else {
			curLength = newLength
		}
	}
	height := max(firstWordsHeight, secondWordsHeight)

	maxSecondWord := 0
	for _, v := range secondWords {
		maxSecondWord = max(maxSecondWord, v)
	}

	l := maxFirstWordsLength
	r := length - maxSecondWord
	for l < r {
		mid := (l + r + 1) / 2
		minHeight, ch := check(firstWords, secondWords, mid, length, height)
		height = minHeight
		if ch {
			l = mid
		} else {
			r = mid - 1
		}
	}

	secondWordsHeight = 1
	curLength = secondWords[0]
	for i, v := range secondWords {
		if i == 0 {
			continue
		}
		newLength := curLength + v + 1
		if newLength > maxSecondWord {
			secondWordsHeight++
			curLength = v
		} else {
			curLength = newLength
		}
	}

	maxFirstWordsLength = length - maxSecondWord

	firstWordsHeight = 1
	curLength = firstWords[0]
	for i, v := range firstWords {
		if i == 0 {
			continue
		}
		newLength := curLength + v + 1
		if newLength > maxFirstWordsLength {
			firstWordsHeight++
			curLength = v
		} else {
			curLength = newLength
		}
	}

	l = maxSecondWord
	r = length - maxFirstWordsLength
	for l < r {
		mid := (l + r + 1) / 2
		minHeight, ch := check(secondWords, firstWords, mid, length, height)
		height = minHeight
		if ch {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return height
}

func check(words []int, foreignWords []int, mid int, length int, height int) (int, bool) {
	wordsLength := mid
	foreignWordsLength := length - mid

	curLength := words[0]
	wordsHeight := 1
	for i, v := range words {
		if i == 0 {
			continue
		}
		newLength := curLength + v + 1
		if newLength > wordsLength {
			wordsHeight++
			curLength = v
		} else {
			curLength = newLength
		}
	}

	curLength = foreignWords[0]
	foreignWordsHeight := 1
	for i, v := range foreignWords {
		if i == 0 {
			continue
		}
		newLength := curLength + v + 1
		if newLength > foreignWordsLength {
			foreignWordsHeight++
			curLength = v
		} else {
			curLength = newLength
		}
	}

	maxHeight := max(wordsHeight, foreignWordsHeight)
	ch := false
	if wordsHeight >= foreignWordsHeight {
		ch = true
	}
	if maxHeight <= height {
		return maxHeight, ch
	} else {
		return height, ch
	}
}
