package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	markArr := make([]bool, cnt)
	markArr[cnt-1] = true
	for i := cnt - 2; i > -1; i-- {
		num := arr[i]
		for j := 1; j < num+1; j++ {
			if i+j < cnt && markArr[i+j] {
				markArr[i] = true
				break
			}
		}
	}

	if markArr[0] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
