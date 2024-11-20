package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	heap.Init()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	blueShirtCnt, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	redShirtCnt, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	blueSockCnt, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	redSockCnt, _ := strconv.Atoi(scanner.Text())

	first := 1
	second := 1

	if blueShirtCnt == 0 && blueSockCnt == 0 {
		fmt.Println("1 1")
		return
	}

	if redShirtCnt == 0 && redSockCnt == 0 {
		fmt.Println("1 1")
		return
	}

	if blueShirtCnt == 0 {
		first = 1
		second = redSockCnt + 1
		fmt.Println(first, second)
		return
	}

	if redShirtCnt == 0 {
		first = 1
		second = blueSockCnt + 1
		fmt.Println(first, second)
		return
	}

	if blueSockCnt == 0 {
		first = redShirtCnt + 1
		second = 1
		fmt.Println(first, second)
		return
	}

	if redSockCnt == 0 {
		first = blueShirtCnt + 1
		second = 1
		fmt.Println(first, second)
		return
	}

	if blueShirtCnt == 1 && redShirtCnt == 1 {
		fmt.Println("2 1")
		return
	}

	if blueSockCnt == 1 && redSockCnt == 1 {
		fmt.Println("1 2")
		return
	}

	first = blueShirtCnt + 1
	second = blueSockCnt + 1
	minCnt := first + second

	tmpFirst := redShirtCnt + 1
	tmpSecond := redSockCnt + 1
	sum := tmpFirst + tmpSecond
	if sum < minCnt {
		first = tmpFirst
		second = tmpSecond
	}

	fmt.Println(first, second)
}
