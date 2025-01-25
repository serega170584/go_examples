package main

import "fmt"

// 10! / (7! * 3!) = 8 * 9 * 10 / (2 * 3) = 4 * 3 * 10 = 120

// c(3) = 2, 2! / (1! * 1!)
// c(4) = 3
// c(5) = 4! / (3! * 1!) * c(3) * c(1) = 4 * 2 * 1 = 8
// c(6) = 5! / (3! * 2!) * c(3) * c(2) = 4 * 5 / 2 * 2 * 1 = 20
// c(7) = 6! / (3! * 3!) * c(3) * c(3) = 4 * 5 * 6 / 6 * 2 * 2 = 80
// c(8) = 7! / (4! * 3!) * c(4) * c(3) = 5 * 6 * 7 / 6 * 3 * 2 = 210
// c(9) = 8! / (5! * 3!) * c(5) * c(3) = 6 * 7 * 8 / 6 * 8 * 2 = 896
// c(10) = 9! / (6! * 3!) * c(6) * c(3) = 7 * 8 * 9 / (2 * 3) * 20 * 2 = 7 * 4 * 3 * 20 * 2 = 3360
// c(11) = 10! / (7! * 3!) * c(7) * c(3) = 8 * 9 * 10 / (2 * 3) * 80 * 2 = 4 * 3 * 10 * 80 * 2 = 19200
// c(12) = 11! / (7! * 4!) * c(7) * c(4) = 8 * 9 * 10 * 11 / (2 * 3 * 4) * 80 * 3 = 3 * 10 * 11 * 80 * 3 = 79200
// c(13) = 12! / (7! * 5!) * c(7) * c(5) = 8 * 9 * 10 * 11 * 12 / (2 * 3 * 4 * 5) * 80 * 8 = 3 * 2 * 11 * 12 * 80 * 8 = 506880
// c(14) = 13! / (7! * 6!) * c(7) * c(6) = 8 * 9 * 10 * 11 * 12 * 13 / (2 * 3 * 4 * 5 * 6) * 80 * 20 = 3 * 2 * 11 * 2 * 13 * 80 * 20 = 2 645 600
// c(20) = 19! / (12! * 7!) * c(12) * c(7) = 13 * 14 * 15 * 16 * 17 * 18 * 19 / (2 * 3 * 4 * 5 * 6 * 7) * 506880 * 80
// 1 - level 1, next level contains 2 in 1
// 2 - level 2, current level contains 2 in 1, 1 from 2 uses
// 3 - level 2, current level contains 2 in 1, 2 from 2 uses, next level contains 2 in 2
func main() {
	fmt.Println(solve(100))
}

func solve(A int) int {
	cnt := make([]int, A+2)
	cnt[1] = 1
	cnt[2] = 1
	if A < 3 {
		return cnt[A]
	}
	levelCnt := 1
	leftLevelCnt := 1
	rightLevelCnt := 0
	koef := 1
	leftCnt := 1
	rightCnt := 0
	for i := 3; i <= A; i++ {
		leftCntChanged := false
		rightCntChanged := false

		if leftLevelCnt == levelCnt {
			rightLevelCnt++
			rightCnt++
			rightCntChanged = true
			if rightLevelCnt == levelCnt {
				levelCnt *= 2
				leftLevelCnt = 0
				rightLevelCnt = 0
			}
		} else {
			leftLevelCnt++
			leftCnt++
			leftCntChanged = true
		}

		koef = koef * (i - 1)
		if leftCntChanged {
			koef /= leftCnt
		}
		if rightCntChanged {
			koef /= rightCnt
		}

		koef = koef % 1000000007

		cnt[i] = (koef * cnt[leftCnt] * cnt[rightCnt]) % 1000000007
	}

	return cnt[A]
}
