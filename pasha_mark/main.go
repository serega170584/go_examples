package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 9 - 3 = 6
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	ind := 1
	participants := make([]string, 2)
	participants[0] = "Pasha"
	participants[1] = "Mark"

	for cnt != 1 {
		delimiter := 2
		differentDelimiter := cnt / delimiter
		mod := cnt % delimiter
		for mod != 0 && delimiter < differentDelimiter {
			mod = cnt % delimiter
			differentDelimiter = cnt / delimiter
			delimiter++
		}
		if mod != 0 {
			break
		}
		cnt -= differentDelimiter
		ind = (ind + 1) % 2
	}

	fmt.Println(participants[ind])
}
