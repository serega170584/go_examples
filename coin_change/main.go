package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	nominals := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		nominals[i], _ = strconv.Atoi(scanner.Text())
	}

	minNominalCnt := make([]int, target+1)
	for i := 1; i < target+1; i++ {
		minNominalCnt[i] = int(math.Inf(1))
	}

	for i := 1; i < target+1; i++ {
		min := int(math.Inf(1))
		for _, nominal := range nominals {
			prevInd := i - nominal
			if prevInd >= 0 && minNominalCnt[prevInd]+1 < min {
				min = minNominalCnt[prevInd] + 1
			}
		}
		minNominalCnt[i] = min
	}

	fmt.Println(minNominalCnt[target])
}
