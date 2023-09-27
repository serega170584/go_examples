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
	cntStr := scanner.Text()
	cnt, _ := strconv.Atoi(cntStr)

	arr := make([]string, 0)
	generate(arr, 0, 0, cnt)
}

func generate(arr []string, open, close, cnt int) {
	if len(arr) == 2*cnt {
		fmt.Println(strings.Join(arr, ""))
	}

	if open != cnt {
		arrCnt := len(arr)
		copyArr := make([]string, arrCnt)
		copy(copyArr, arr)
		copyArr = append(copyArr, "(")
		generate(copyArr, open+1, close, cnt)
	}

	if open > close {
		arrCnt := len(arr)
		copyArr := make([]string, arrCnt)
		copy(copyArr, arr)
		copyArr = append(copyArr, ")")
		generate(copyArr, open, close+1, cnt)
	}
}
