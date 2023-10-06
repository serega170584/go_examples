package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// abba
// aba
// abacaba
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	arr := make([]string, cnt)

	for i := 0; i < cnt; i++ {
		scanner.Scan()
		arr[i] = scanner.Text()
	}

	maxSubstring := make([]string, cnt)
	maxLen := 0

	for i := range arr {

		substring := make([]string, cnt)
		subLen := 0

		left, right := i, i

		for left != -1 && right != cnt && arr[left] == arr[right] {
			substring[left] = arr[left]
			substring[right] = arr[right]

			subLen++
			left--
			right++
		}

		if subLen > maxLen {
			maxSubstring = substring
			maxLen = subLen
		}
	}

	for i := range arr {

		substring := make([]string, cnt)
		subLen := 0

		left, right := i, i+1

		for left != -1 && right != cnt && arr[left] == arr[right] {
			substring[left] = arr[left]
			substring[right] = arr[right]

			subLen++
			left--
			right++
		}

		if subLen > maxLen {
			maxSubstring = substring
			maxLen = subLen
		}
	}

	fmt.Println(maxSubstring)
}
