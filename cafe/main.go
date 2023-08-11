package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// 5 35 40 101 59 63
// 6 35 40 101 59 63 101
// 4 35 40 101 0
// 7 35 40 101 59 63 101 40
// 7 35 40 101 59 53 101 40
// 8 35 40 101 59 53 101 70 80
// 8 35 40 101 90 53 101 70 80
// 1 35
func main() {
	var cnt int
	_, err := fmt.Scanln(&cnt)
	if err != nil {
		log.Fatal(err)
	}

	dinners := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		_, err = fmt.Scanln(&dinners[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	baseFreeDinnerIndexes := make([]int, 0)
	secondaryFreeDinnerIndexes := make([]int, 0)

	minBaseFreeDinnerIndex := -1
	minSecondaryFreeDinnerIndex := -1

	minBaseFreeDinnerIndexOrder := 0
	minSecondaryFreeDinnerIndexOrder := 0

	var priceSum, freeDinnerSum, secondaryFreeDinnerSum, expensiveDinnerCnt, startInd, notUsedCouponsCnt, secondaryNotUsedCouponsCnt, usedCouponsCnt, secondaryUsedCouponsCnt int

	resultFreeIndexes := make([]string, 0)
	resultSecondaryFreeIndexes := make([]string, 0)

	for ind, dinnerPrice := range dinners {
		priceSum += dinnerPrice
		startInd = ind + 1

		if dinnerPrice == 0 {
			continue
		}

		if dinnerPrice < 101 && expensiveDinnerCnt == 1 && (minBaseFreeDinnerIndex == -1 || dinners[minBaseFreeDinnerIndex] < dinnerPrice) {
			baseFreeDinnerIndexes[0] = ind
			secondaryFreeDinnerIndexes[0] = ind
			minBaseFreeDinnerIndex = ind
			continue
		}

		if dinnerPrice > 100 {
			expensiveDinnerCnt++
		}

		if dinnerPrice > 100 && expensiveDinnerCnt == 1 {
			baseFreeDinnerIndexes = append(baseFreeDinnerIndexes, -1)
			secondaryFreeDinnerIndexes = append(secondaryFreeDinnerIndexes, -1)
		}

		if expensiveDinnerCnt == 2 {
			baseFreeDinnerIndexes = append(baseFreeDinnerIndexes, -1)
			secondaryFreeDinnerIndexes[0] = ind
			minSecondaryFreeDinnerIndex = ind
			break
		}
	}

	prevInd := -1
	for ind := startInd; ind < cnt; ind++ {
		priceSum += dinners[ind]

		if prevInd == -1 {
			baseFreeDinnerIndexes = append(baseFreeDinnerIndexes, ind)
		}

		if prevInd == -1 && minBaseFreeDinnerIndex != -1 && dinners[ind] < dinners[minBaseFreeDinnerIndex] {
			minBaseFreeDinnerIndex = ind
			minBaseFreeDinnerIndexOrder++
		}

		if prevInd != -1 && minBaseFreeDinnerIndex != -1 && dinners[minBaseFreeDinnerIndex] < dinners[ind] {
			baseFreeDinnerIndexes[minBaseFreeDinnerIndexOrder] = ind
			for order, baseFreeDinnerIndex := range baseFreeDinnerIndexes {
				if baseFreeDinnerIndex == -1 {
					minBaseFreeDinnerIndexOrder = order
					minBaseFreeDinnerIndex = -1
				} else if baseFreeDinnerIndex < minBaseFreeDinnerIndex {
					minBaseFreeDinnerIndexOrder = order
					minBaseFreeDinnerIndex = baseFreeDinnerIndex
				}
			}
		}

		if dinners[ind] > 100 {
			if dinners[minSecondaryFreeDinnerIndex] < dinners[ind] {
				for order, secondaryFreeDinnerIndex := range secondaryFreeDinnerIndexes {
					if secondaryFreeDinnerIndex == -1 {
						minSecondaryFreeDinnerIndexOrder = order
						minSecondaryFreeDinnerIndex = -1
					} else if secondaryFreeDinnerIndex < minSecondaryFreeDinnerIndex {
						minSecondaryFreeDinnerIndexOrder = order
						minSecondaryFreeDinnerIndex = secondaryFreeDinnerIndex
					}
				}
			}

			baseFreeDinnerIndexes, secondaryFreeDinnerIndexes = secondaryFreeDinnerIndexes, baseFreeDinnerIndexes
			minBaseFreeDinnerIndexOrder, minSecondaryFreeDinnerIndexOrder = minSecondaryFreeDinnerIndexOrder, minBaseFreeDinnerIndexOrder
			minBaseFreeDinnerIndex, minSecondaryFreeDinnerIndex = minSecondaryFreeDinnerIndex, minBaseFreeDinnerIndex

			prevInd = -1

			break
		}

		prevInd = ind
	}

	for _, baseFreeDinnerIndex := range baseFreeDinnerIndexes {
		if baseFreeDinnerIndex == -1 {
			notUsedCouponsCnt++
		} else {
			freeDinnerSum += dinners[baseFreeDinnerIndex]
			usedCouponsCnt++
			resultFreeIndexes = append(resultFreeIndexes, strconv.Itoa(baseFreeDinnerIndex+1))
		}
	}

	for _, secondaryFreeDinnerIndex := range secondaryFreeDinnerIndexes {
		if secondaryFreeDinnerIndex == -1 {
			secondaryNotUsedCouponsCnt++
		} else {
			secondaryFreeDinnerSum += dinners[secondaryFreeDinnerIndex]
			secondaryUsedCouponsCnt++
			resultSecondaryFreeIndexes = append(resultSecondaryFreeIndexes, strconv.Itoa(secondaryFreeDinnerIndex+1))
		}
	}

	if freeDinnerSum < secondaryFreeDinnerSum {
		freeDinnerSum = secondaryFreeDinnerSum
		notUsedCouponsCnt = secondaryNotUsedCouponsCnt
		usedCouponsCnt = secondaryUsedCouponsCnt
		resultFreeIndexes = resultSecondaryFreeIndexes
	}

	fmt.Println(priceSum - freeDinnerSum)
	fmt.Printf("%d %d\n", notUsedCouponsCnt, usedCouponsCnt)
	fmt.Println(strings.Join(resultFreeIndexes, " "))

}
