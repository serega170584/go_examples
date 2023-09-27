package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	s1 := scanner.Text()

	scanner.Scan()
	s2 := scanner.Text()

	dict1 := make(map[int32]int, len(s1))
	dict2 := make(map[int32]int, len(s2))

	valList := make([]int32, len(s1)+len(s2))
	var valListPointer int

	for _, val := range s1 {
		if _, ok := dict1[val]; !ok {
			dict1[val] = 0
			valList[valListPointer] = val
			valListPointer++
		}
		dict1[val]++
	}

	for _, val := range s2 {
		if _, ok := dict2[val]; !ok {
			dict2[val] = 0
			valList[valListPointer] = val
			valListPointer++
		}
		dict2[val]++
	}

	isNo := false
	for _, val := range valList {
		if val == 0 {
			break
		}

		if _, ok := dict1[val]; !ok {
			fmt.Println("NO")
			isNo = true
			break
		}

		if _, ok := dict2[val]; !ok {
			fmt.Println("NO")
			isNo = true
			break
		}

		if dict1[val] != dict2[val] {
			fmt.Println("NO")
			isNo = true
			break
		}
	}

	if !isNo {
		fmt.Println("YES")
	}
}
