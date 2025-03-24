package main

import (
	"fmt"
)

func main() {
	//a := []int{11, 10, 8, 6, 8, 11, 1, 10, 2, 3, 8, 3, 8, 12, 11, 1, 7, 5, 5, 12, 9, 4, 10, 3, 3, 3, 8, 8, 8, 6, 7, 7, 7, 6, 4, 2, 5, 8, 11, 10, 10, 10, 12, 9, 2, 3, 9, 12, 7, 6, 11, 8, 9, 9, 10, 3, 3, 5, 2, 10, 10, 9, 4, 9, 6, 11, 10, 2, 6, 1, 4, 7, 10, 3, 4, 3, 9, 4, 3, 8, 1, 1, 3}
	//a := []int{2, 1, 3, 3, 5, 1, 4, 3, 5, 3, 1, 1, 4, 5, 1, 4, 3, 7, 5, 3, 6, 3, 5, 5, 4, 7, 1, 3, 3, 4, 2, 4, 5, 3, 3, 5, 6, 2, 2, 7, 6, 2, 4, 4, 4, 1, 5, 5, 7, 2, 7, 6, 3, 2, 7, 1, 7, 5, 1, 4, 2, 5, 8, 8, 8, 8, 8, 10, 9, 10, 9, 9, 10, 9, 10, 9, 8, 10, 9, 8, 9, 9, 9, 10, 10, 10, 10, 9, 10, 9, 10}
	a := []int{14, 10, 4}
	// 10 - 11
	// 9 - 11
	// 8 - 7
	// 110 + 99 + 56 = 265
	// 110 + 99 + 40 = 249
	//a := []int{3, 3, 7, 10}
	sum := 0
	l := len(a)
	fmt.Println(l)
	for i := 0; i < l; i++ {
		sum += a[i]
	}

	fmt.Println(sum)

	halfSum := sum/2 + sum%2
	fmt.Println(halfSum)

	dp := make([][]int, l)
	dp[0] = make([]int, halfSum)

	amount := make([][]int, l)
	amount[0] = make([]int, halfSum)

	for i := 0; i < halfSum; i++ {
		if i+1 >= a[0] {
			dp[0][i] = a[0]
			amount[0][i] = 1
		}
	}

	for i := 1; i < l; i++ {
		dp[i] = make([]int, halfSum)
		amount[i] = make([]int, halfSum)
		for j := 0; j < halfSum; j++ {
			if j+1-a[i] == 0 {
				oldVal := dp[i-1][j]
				oldAmount := amount[i-1][j]
				newVal := a[i]
				newAmount := 1
				dp[i][j] = oldVal
				amount[i][j] = oldAmount
				if newVal >= oldVal {
					if newVal == oldVal {
						if newAmount > oldAmount {
							dp[i][j] = newVal
							amount[i][j] = newAmount
						}
					} else {
						dp[i][j] = newVal
						amount[i][j] = newAmount
					}
				}
				continue
			}
			if j-a[i] >= 0 {
				oldVal := dp[i-1][j]
				oldAmount := amount[i-1][j]
				newVal := dp[i-1][j-a[i]] + a[i]
				newAmount := amount[i-1][j-a[i]] + 1
				dp[i][j] = oldVal
				amount[i][j] = oldAmount
				if newVal >= oldVal {
					if newVal == oldVal {
						if newAmount > oldAmount {
							dp[i][j] = newVal
							amount[i][j] = newAmount
						}
					} else {
						dp[i][j] = newVal
						amount[i][j] = newAmount
					}
				}
			} else {
				dp[i][j] = dp[i-1][j]
				amount[i][j] = amount[i-1][j]
			}
		}
		fmt.Println("a[i]", a[i])
		fmt.Println("amount: ", amount[i][halfSum-1])
		fmt.Println("sum: ", dp[i][halfSum-1])
	}

	//fmt.Println(dp)
	//fmt.Println(amount)

	fmt.Println(dp[l-1][halfSum-1])

	fmt.Println(l - amount[l-1][halfSum-1])

}
