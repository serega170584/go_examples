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

	list := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(canJump(list))
}

func canJump(A []int) int {
	if A[0] == 0 && len(A) == 1 {
		return 1
	}

	maxReach := 0
	currReach := 0
	jump := 0
	for i := 0; i < len(A); i++ {
		if i+A[i] > maxReach {
			maxReach = i + A[i]
		}

		if maxReach >= len(A)-1 {
			jump++
			break
		}

		if i == currReach {
			if i == maxReach {
				return 0
			}

			currReach = maxReach
			jump++
		}
	}

	return jump
}
