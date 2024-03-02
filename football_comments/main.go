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
	firstCount := scanner.Text()

	scanner.Scan()
	secondCount := scanner.Text()

	scanner.Scan()
	firstPlay, _ := strconv.Atoi(scanner.Text())

	fmt.Println(getBallQuantity(firstCount, secondCount, firstPlay))
}

func getBallQuantity(firstCount string, secondCount string, firstPlay int) int {
	firstRecords := strings.Split(firstCount, ":")
	firstBeats, _ := strconv.Atoi(firstRecords[0])
	firstLost, _ := strconv.Atoi(firstRecords[1])

	secondRecords := strings.Split(secondCount, ":")
	secondBeats, _ := strconv.Atoi(secondRecords[0])
	secondLost, _ := strconv.Atoi(secondRecords[1])

	balance := firstBeats - firstLost + secondBeats - secondLost
	if balance > 0 {
		return 0
	}

	if balance == 0 {
		if firstPlay == 1 {
			if firstLost < secondBeats {
				return 0
			} else {
				return 1
			}
		}

		if firstPlay == 2 {
			if secondLost < firstBeats {
				return 0
			} else {
				return 1
			}
		}
	}

	leftBeats := -balance
	if firstPlay == 1 {
		if firstLost < secondBeats+leftBeats {
			return leftBeats
		} else {
			return leftBeats + 1
		}
	}

	if firstPlay == 2 {
		if secondLost < firstBeats {
			return leftBeats
		}
	}

	return leftBeats + 1
}
