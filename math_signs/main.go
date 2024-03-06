package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	list := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(getSigns(cnt, list))
}

func getSigns(cnt int, list []int) string {
	var sb strings.Builder

	isOdd := false
	if list[0]%2 != 0 {
		isOdd = true
	}
	for i := 1; i < cnt; i++ {
		if isOdd {
			if list[i]%2 == 0 {
				sb.WriteString(string(rune(43)))
			} else {
				sb.WriteString(string(rune(120)))
			}
		} else {
			sb.WriteString(string(rune(43)))
			isOdd = false
			if list[i]%2 != 0 {
				isOdd = true
			}
		}
	}

	return sb.String()
}
