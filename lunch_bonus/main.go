package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Node struct {
	prevI int
	prevJ int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter days count")
	scanner.Scan()
	daysCnt, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list")
	list := make([]int, daysCnt)
	for i := 0; i < daysCnt; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	cost, path := minCost(daysCnt, list)

	fmt.Println("Got min cost", cost)
	fmt.Println("Got min path")

	for _, v := range path {
		fmt.Println(*v)
	}
}

func minCost(daysCnt int, list []int) (int, []*Node) {
	couponsCnt := 0
	for _, v := range list {
		if v > 100 {
			couponsCnt++
		}
	}

	dp := make([][]int, couponsCnt+1)
	path := make([][]*Node, couponsCnt+1)
	for i := 0; i < couponsCnt+1; i++ {
		dp[i] = make([]int, daysCnt)
		path[i] = make([]*Node, daysCnt)
		for j := 0; j < daysCnt; j++ {
			dp[i][j] = math.MaxInt
		}
	}

	dp[0][0] = list[0]
	for j := 0; j < daysCnt; j++ {
		for i := 0; i < couponsCnt+1; i++ {
			nodes := make(map[int]*Node, 2)

			prevDay := j - 1
			nextCoupon := i + 1
			coupon := i
			if list[j] > 100 {
				coupon = i - 1
			}
			prevDayWithCoupon := math.MaxInt
			prevDayWithoutCoupon := math.MaxInt

			if prevDay >= 0 && coupon >= 0 && dp[coupon][prevDay] != math.MaxInt {
				prevDayWithoutCoupon = dp[coupon][prevDay] + list[j]
				nodes[prevDayWithoutCoupon] = &Node{prevI: coupon, prevJ: prevDay}
			}

			if prevDay >= 0 && nextCoupon <= couponsCnt && dp[nextCoupon][prevDay] != math.MaxInt {
				prevDayWithCoupon = dp[nextCoupon][prevDay]
				nodes[prevDayWithCoupon] = &Node{prevI: nextCoupon, prevJ: prevDay}
			}
			dp[i][j] = min(dp[i][j], prevDayWithoutCoupon, prevDayWithCoupon)
			path[i][j] = nodes[dp[i][j]]
		}
	}

	minVals := make([]int, couponsCnt+1)
	pathList := make([]*Node, daysCnt)
	nodes := make(map[int]*Node, couponsCnt+1)
	for i := 0; i < couponsCnt+1; i++ {
		prevI := i - 1
		minVals[i] = dp[i][daysCnt-1]
		nodes[dp[i][daysCnt-1]] = &Node{prevI: i, prevJ: daysCnt - 1}
		if prevI >= 0 {
			minVals[i] = min(minVals[i], minVals[i-1])
		}
	}
	pathList[daysCnt-1] = nodes[minVals[couponsCnt]]
	for i := daysCnt - 2; i >= 0; i-- {
		pathI := pathList[i+1].prevI
		pathJ := pathList[i+1].prevJ
		pathList[i] = path[pathI][pathJ]
	}
	return minVals[couponsCnt], pathList
}
