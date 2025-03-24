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
	firstStart, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	firstInterval, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	secondStart, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	secondInterval, _ := strconv.Atoi(scanner.Text())

	leastLeft := firstStart - firstInterval
	leastRight := firstStart + firstInterval

	mostLeft := secondStart - secondInterval
	mostRight := secondStart + secondInterval

	if leastLeft > mostLeft {
		leastLeft, mostLeft = mostLeft, leastLeft
		leastRight, mostRight = mostRight, leastRight
	}

	s := 0
	if mostLeft > leastRight {
		s += leastRight - leastLeft + 1
		s += mostRight - mostLeft + 1
	} else {
		s += mostRight - leastLeft + 1
	}

	fmt.Println(s)
}
