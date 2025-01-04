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
	firstPoint, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	firstDistance, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	secondPoint, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	secondDistance, _ := strconv.Atoi(scanner.Text())

	firstLeft := firstPoint - firstDistance
	firstRight := firstPoint + firstDistance

	distance := firstRight - firstLeft + 1

	secondLeft := secondPoint - secondDistance
	secondRight := secondPoint + secondDistance

	if secondLeft > firstRight || firstLeft > secondRight {
		distance += secondRight - secondLeft + 1
		fmt.Println(distance)
		return
	}

	if firstLeft >= secondLeft {
		distance += firstLeft - secondLeft
	}

	if firstRight <= secondRight {
		distance += secondRight - firstRight
	}

	fmt.Println(distance)
}
