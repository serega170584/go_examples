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

	cntStr := scanner.Text()
	cnt, _ := strconv.Atoi(cntStr)

	arr := make([]string, cnt+1)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		arr[i] = scanner.Text()
	}

	startPointer := 0
	finishPointer := 0
	for finishPointer < cnt+1 {
		for arr[startPointer] == arr[finishPointer] {
			finishPointer++
		}

		if finishPointer == cnt {
			break
		}

		startPointer++
		arr[startPointer] = arr[finishPointer]
		finishPointer++
	}

	for i := 0; i <= startPointer; i++ {
		fmt.Println(arr[i])
	}

}
