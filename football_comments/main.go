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
	firstScore := scanner.Text()

	scanner.Scan()
	secondScore := scanner.Text()

	scanner.Scan()
	host, _ := strconv.Atoi(scanner.Text())

	parts := strings.Split(firstScore, ":")
	hostBeat, _ := strconv.Atoi(parts[0])
	hostLost, _ := strconv.Atoi(parts[1])
	parts = strings.Split(secondScore, ":")
	guestBeat, _ := strconv.Atoi(parts[0])
	guestLost, _ := strconv.Atoi(parts[1])
	if host == 2 {
		parts = strings.Split(secondScore, ":")
		hostBeat, _ = strconv.Atoi(parts[0])
		hostLost, _ = strconv.Atoi(parts[1])
		parts = strings.Split(firstScore, ":")
		guestBeat, _ = strconv.Atoi(parts[0])
		guestLost, _ = strconv.Atoi(parts[1])
	}

	diff := hostBeat - hostLost + guestBeat - guestLost
	if diff > 0 || (diff == 0 && guestBeat > hostLost) {
		fmt.Println("0")
		return
	}

	rest := -diff
	if host == 1 && guestBeat+rest <= hostLost {
		rest++
	}

	if host == 2 && guestBeat <= hostLost {
		rest++
	}

	fmt.Println(rest)

}
