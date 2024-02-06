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

	fmt.Println("Enter list count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	byteScanner := bufio.NewScanner(os.Stdin)
	byteScanner.Split(bufio.ScanBytes)

	fmt.Println("Enter list")
	list := make([]byte, n)
	for i := 0; i < n; i++ {
		byteScanner.Scan()
		list[i] = byteScanner.Bytes()[0]
	}

	fmt.Println(getRLE(n, list))
}

func getRLE(n int, list []byte) string {
	var sb strings.Builder
	var prev byte
	counter := 0
	for i, v := range list {
		if prev == 0 {
			prev = v
			counter++
			continue
		}
		if prev != v {
			sb.WriteByte(prev)
			sb.WriteString(strconv.Itoa(counter))
			prev = v
			counter = 0
		}
		counter++
		if i == n-1 {
			sb.WriteByte(prev)
			sb.WriteString(strconv.Itoa(counter))
		}
	}
	return sb.String()
}
