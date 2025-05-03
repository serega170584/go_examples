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
	a, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	d, _ := strconv.Atoi(scanner.Text())

	fmt.Println(smallestSequenceWithPrimes(a, b, c, d))
}

func smallestSequenceWithPrimes(A int, B int, C int, D int) []int {
	if D == 0 {
		return []int{}
	}

	if D == 1 {
		return []int{A}
	}

	primes := []int{A, B, C}

	dp := make([]int, 1, D+1)
	dp[0] = 1

	indexes := make([]int, 3)

	smallestSequence := make([]int, 0, D)

	existed := make(map[int]struct{})

	for D != 0 {
		minPrimeIndex := 0
		for i := 1; i < 3; i++ {
			if dp[indexes[i]]*primes[i] < dp[indexes[minPrimeIndex]]*primes[minPrimeIndex] {
				minPrimeIndex = i
			}
		}

		a := dp[indexes[minPrimeIndex]]
		b := primes[minPrimeIndex]
		val := a * b
		if _, ok := existed[val]; ok {
			indexes[minPrimeIndex]++
			continue
		}
		existed[dp[indexes[minPrimeIndex]]*primes[minPrimeIndex]] = struct{}{}

		dp = append(dp, dp[indexes[minPrimeIndex]]*primes[minPrimeIndex])
		smallestSequence = append(smallestSequence, dp[indexes[minPrimeIndex]]*primes[minPrimeIndex])
		D--
	}

	return smallestSequence
}
