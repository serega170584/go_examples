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
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())

	fmt.Println(digitNumbersWithDigitSum(n, s))
}

func digitNumbersWithDigitSum(A int, B int) int {
	sumItems := []int{0}
	digits := [][]int{{}}
	n := A

	curSumItems := []int{}
	curDigits := [][]int{}
	for i := 0; i < len(sumItems); i++ {
		for j := 1; j <= 9; j++ {
			v := sumItems[i] + j
			if v <= B {
				curSumItems = append(curSumItems, v)
				curDigits = append(curDigits, append(digits[i], j))
			}
		}
	}
	sumItems = curSumItems
	digits = curDigits
	n--

	for n != 0 {
		curSumItems = []int{}
		curDigits = [][]int{}
		for i := 0; i < len(sumItems); i++ {
			for j := 0; j <= 9; j++ {
				v := sumItems[i] + j
				if v <= B {
					curSumItems = append(curSumItems, v)
					curDigits = append(curDigits, append(digits[i], j))
				}
			}
		}
		sumItems = curSumItems
		digits = curDigits
		n--
	}

	numberDigits := []int{}
	for i := 0; i < len(curSumItems); i++ {
		if curSumItems[i] == B {
			numberDigits = digits[i]
			break
		}
	}

	mod := numberDigits[len(numberDigits)-1]
	mult := 10
	for i := len(numberDigits) - 2; i >= 0; i-- {
		mod = ((numberDigits[i]*mult)%1000000007 + mod) % 1000000007
		mult = (mult * 10) % 1000000007
	}

	return mod
}
