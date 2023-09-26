package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var j, s string

	scanner.Scan()
	j = scanner.Text()

	scanner.Scan()
	s = scanner.Text()

	dict := make(map[string]struct{}, len(j))
	for _, val := range j {
		str := fmt.Sprintf("%c", val)
		dict[str] = struct{}{}
	}

	var cnt int
	for _, val := range s {
		str := fmt.Sprintf("%c", val)
		if _, ok := dict[str]; ok {
			cnt++
		}
	}

	fmt.Println(cnt)

}
