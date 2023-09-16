package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 6 2 4
// 2 + 4 = 6
// 6 + 6 = 12
// 6 + 12 = 18

// 1 2 3
// 1 + 2 = 3
// 3 3
// 3 + 3 = 6
// 3 + 6  = 9

// 1 2 3 4
// 1 + 2 = 3
// 3 3 4
// 3 + 3 = 6
// 4 6
// 6 + 4 = 10
// 3 + 6 + 10  = 19

// 4 3 2
// 2 + 3 = 5
// 4 5
// 4 + 5 = 9
// 5 + 9 = 14

// 4 3 2 1
// 2 + 3 = 5
// 4 5 1
// 1 + 5 = 6
// 5 5
// 5 + 5 = 10
// 5 + 6 + 10 = 21

// 2 + 3
// 1 enter
// 1 + 2
// 1 + 2 3

// 2 3
// 2
// 2 3

// 2 3 1
// 1
// 2 1
// 2 3 1

// 2 3 1 4
// 1
// 2 1
// 2 3 1
// 2 3 1 4

// 2 5 8 11 3 7 4 1

// 2
// 2

// 2 5
// 2
// 2 5

// 2 5 8
// 2
// 2 5
// 2 5 8

// 2 5 8 11
// 2
// 2 5
// 2 5 8
// 2 5 8 11

// 2 5 8 11 3
// 2
// 2 3
// 2 5 3
// 2 5 8 3
// 2 5 8 11 3

// 2 5 8 11 3 7
// 2
// 2 3
// 2 5 3
// 2 5 3 7
// 2 5 8 3 7
// 2 5 8 11 3 7

// 2 5 8 11 3 7 4
// 2
// 2 3
// 2 3 4
// 2 5 3 4
// 2 5 3 7 4
// 2 5 8 3 7 4
// 2 5 8 11 3 7 4

// 2 5 8 11 3 7 4 1
// 1
// 2 1
// 2 3 1
// 2 3 4 1
// 2 5 3 4 1
// 2 5 3 7 4 1
// 2 5 8 3 7 4 1
// 2 5 8 11 3 7 4

// 6 2 4 1
// 2 + 1 = 3
// 3 + 4 = 7
// 7 + 6 = 13
// 3 + 7 + 13 = 23

// 2 5 8 11 3 7 4 1
// 1 + 2 = 3
// 3 + 3 = 6
// 6 + 4 = 10
// 10 + 5 = 15
// 15 + 7 = 22
// 22 + 8 = 30
// 30 + 11 = 41
// 3 + 6 + 10 + 15 + 22 + 30 + 41 = 127

// 09:44 - 11:17
// 1. count = 1 then output 0 index element
// 2. count = 2 then output sum of 0 and 1 element
// 3. create array of states: sum for next elem: prev sum + current number or current state element
// variables:
// 1. cnt
// 2. array brick's counts

// 2 1 1 1 1 1 1 1 1 1
// 1 + 1 = 2
// 2 + 1 = 3
// 3 + 1 = 4
// 4 + 1 = 5
// 5 + 1 = 6
// 6 + 1 = 7
// 7 + 1 = 8
// 8 + 1 = 9
// 9 + 2 = 11

// 2 1 1 1 1 1 1 1 1 1
// 1 + 1 = 2
// 2 1 1 1 1 1 1 1 2
// 1 + 1 = 2
// 2 1 1 1 1 1 2 2
// 1 + 1 = 2
// 2 1 1 1 2 2 2
// 1 + 1 = 2
// 2 1 2 2 2 2
// 1 + 2 = 3
// 2 2 2 2 3
// 2 + 2 = 4
// 2 2 4 3
// 2 + 2 = 4
// 4 4 3
// 3 + 4 = 7
// 4 7
// 4 + 7 = 11
// 2 + 2 + 2 + 2 + 3 + 4 + 4 + 7 + 11 = 37

// 13:06 - 14:20
// 14:39 - 15:22
// 7 4 6 9
// 7 4
// 4
// 7 4

//7 4 6
// 4
// 4 6
// 7 4 6

//7 4 6 9
// 4
// 4 6
// 7 9
// 7 4 6 9

// 2 1 1 1 1 1 1 1 1 1
// 1
// 2 1

// 1
// 1 1
// 2 1 1

// 1
// 1 1
// 1 1 1
// 2 1 1 1

// 1
// 1 1
// 1 1 1
// 1 1 1 1
// 2 1 1 1 1

// 16:45 - 16:58
// 19:15 -

// 2 1 1 1 1 1 1 1 1 1
// 1 + 1
// 1 + 1
// 1 + 1
// 1 + 1
// 2 + 1
// 2 + 2
// 2 + 2
// 3 + 4
// 7 + 4
// 2 + 2 + 2 + 2 + 3 + 4 + 4 + 7 + 11 = 37

// 6 2 5 9

// 6 2
// 2
// 6 + 2 = 8

// 6 2 5
// 2
// 2 + 5 = 7
// 6 + 7 = 13

// 6 2 5 9
// 2
// 2 + 5
// 6
// 6 + 9
// 2 + 5 + 6 + 9

// 6 2 5 9 3
// 2
// 2 + 3 = 5
// 2 + 3 + 5 = 10
// 6
// 6 + 9 = 15
// 2 + 3 + 5 + 6 + 9
// 5 + 10 + 15

// 2 + 3 + 5 and 6 + 9
// 10 + 15 = 25

// 2 + 3 + 5 = 10

// 6 2 5 9 3 1
// 1
// 1 + 2
// 1 + 2 + 3
// 5
// 5 + 6
// 1 + 2 + 3 + 9
// 1 + 2 + 3 + 5 + 6 + 9

// 6 2 5 9 3 1 2
// 1
// 1 + 2
// 2
// 2 + 3
// 1 + 2 + 5
// 1 + 2 + 5 + 6
// 1 + 2 + 2 + 3 + 5 + 6 + 9

// 6 2
// 8

// 6 2 4
// 18

// 6 2 4 5
// 2 + 4 = 6
// 5 + 6 = 11
// 6 + 11 = 17
// 34
// 34 - 18 = 16

// 1 2 3 4 5 6 7 8 9
// 1 + 2 = 3
// 3 + 3 = 6
// 4 + 5 = 9
// 6 + 7 = 13

// 3 4 5 6 10
// 3 + 4 = 7
// 5 + 6 = 11

// 1 1 1 1
// 1 + 1
// 1 + 1
// 2 + 2

// 2 1 1 1 1 1 1 1 1 1
// 2 + 1 = 3
// 1 + 1 = 2
// 1 + 1 = 2
// 1 + 1 = 2
// 1 + 1 = 2
// 2 + 2 = 4
// 2 + 2 = 4
// 3 + 4 = 7
// 7 + 4 = 11
// 3 + 2 + 2 + 2 + 2 + 4 + 4 + 7 + 11 = 37

// 1 1 1 1
// 1 + 1
// 1 + 1
// 2 + 2

// 1 1 1
// 1 + 1 = 2
// 2 + 1 = 3

// 2 4 6
// 6 12

// 6 12
// 18

// 1 1 1 1
// 2 2
// 4

// 1 1 1
// 2

// 1 1 1 1
// 2 2
// 4

// 1 1 1 1 1
// 2 2 3
// 4 7
// 11

func main() {
	var cnt int
	_, err := fmt.Scan(&cnt)
	if err != nil {
		log.Fatal(err)
	}

	scanner := makeScanner()
	stones := readArray(scanner)
	sort.Ints(stones)

	fmt.Println(getEnergyForUnion(stones, cnt))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for ind, val := range listString {
		arr[ind], _ = strconv.Atoi(val)
	}
	return arr
}

func getEnergyForUnion(stones []int, cnt int) int {
	if cnt == 1 {
		return 0
	}

	var min int

	cornerSumInd := cnt - 1

	termInd := 1

	sumInd := -1

	minSumInd := -1

	prevTermInd := 0

	energySum := make([]int, cnt)

	for cornerSumInd != 0 {
		if sumInd == -1 {
			sum := stones[termInd] + stones[termInd-1]
			min += sum
			sumInd++
			minSumInd++
			prevTermInd = termInd
			energySum[sumInd] = sum
			termInd += 2
			cornerSumInd--
			continue
		}

		if termInd == cnt && termInd-prevTermInd == 2 {
			prevTermInd++
		}

		if termInd == cnt {
			sum := stones[prevTermInd] + energySum[minSumInd]
			min += sum
			sumInd++
			energySum[sumInd] = sum

			cnt = sumInd + 1

			cornerSumInd--

			stones, energySum = energySum, stones

			prevTermInd = 0
			termInd = 1
			sumInd = -1
			minSumInd = -1
			continue
		}

		if termInd == cnt+1 {
			stones, energySum = energySum, stones

			cnt = sumInd + 1

			prevTermInd = 0
			termInd = minSumInd + 1
			sumInd = -1
			minSumInd = -1
			continue
		}

		prevTerm := stones[termInd-1]

		if energySum[minSumInd]+prevTerm > prevTerm+stones[termInd] {
			sum := prevTerm + stones[termInd]
			min += sum
			sumInd++
			energySum[sumInd] = sum
			prevTermInd = termInd
			termInd += 2
			cornerSumInd--
			continue
		}

		if energySum[minSumInd]+prevTerm <= prevTerm+stones[termInd] {
			sum := energySum[minSumInd] + prevTerm
			min += sum
			sumInd++
			minSumInd++
			energySum[sumInd] = sum
			prevTermInd = termInd
			termInd++
			cornerSumInd--
			continue
		}
	}

	return min
}
