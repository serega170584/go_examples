package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// 16 = 2 * 2 * 2 * 2
// 8 = 2 * 2 * 2
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter number")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter max used number")
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Got min signs count:", symsMin(k, n))
}

func symsMin(k int, n int) int {
	sumMinList := make([]int, n+1)
	multMinList := make([]int, n+1)
	for i := 0; i < k+1; i++ {
		sumMinList[i] = 1
		multMinList[i] = 1
	}

	for i := k + 1; i < n+1; i++ {
		sumMinList[i] = math.MaxInt
		multMinList[i] = math.MaxInt
	}

	for i := k + 1; i < n+1; i++ {
		for j := 1; j < i; j++ {
			firstSummand := min(multMinList[j], sumMinList[j])
			secondSummand := min(multMinList[i-j], sumMinList[i-j])
			sumMinList[i] = firstSummand + secondSummand + 1

			mod := i % j
			if mod == 0 && j != 1 {
				firstMultiplier := min(multMinList[j], sumMinList[j]+2)
				secondMultiplier := min(multMinList[i/j], sumMinList[i/j]+2)
				multMinList[i] = firstMultiplier + secondMultiplier + 1
			}
		}
	}

	fmt.Println(sumMinList)
	fmt.Println(multMinList)

	return min(sumMinList[n], multMinList[n])
}
