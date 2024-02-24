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

	fmt.Println("Enter shops count")
	scanner.Scan()
	shopsCnt, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter whole sale amount")
	wholeSaleAmounts := make([]int, shopsCnt)
	for i := 0; i < shopsCnt; i++ {
		scanner.Scan()
		wholeSaleAmounts[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Enter whole sale price")
	wholeSalePrice := make([]int, shopsCnt)
	for i := 0; i < shopsCnt; i++ {
		scanner.Scan()
		wholeSalePrice[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Enter retail price")
	retailPrice := make([]int, shopsCnt)
	for i := 0; i < shopsCnt; i++ {
		scanner.Scan()
		retailPrice[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Enter bricks count")
	scanner.Scan()
	bricksCnt, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Got min price", minBrickPrice(shopsCnt, bricksCnt, wholeSaleAmounts, wholeSalePrice, retailPrice))
}

func minBrickPrice(shopsCnt int, bricksCnt int, wholeSaleAmount []int, wholeSalePrice []int, retailPrice []int) int {
	maxWholeSaleAmount := 0
	for _, v := range wholeSaleAmount {
		maxWholeSaleAmount = max(maxWholeSaleAmount, v+2)
	}
	cnt := max(maxWholeSaleAmount, bricksCnt+1)

	dp := make([][]int, shopsCnt)

	for i := 0; i < shopsCnt; i++ {
		dp[i] = make([]int, cnt)

		for j := 0; j < cnt; j++ {
			dp[i][j] = math.MaxInt
		}

		dp[i][0] = 0
	}

	for j := 0; j < cnt; j++ {
		dp[0][j] = j * retailPrice[0]
		if j > wholeSaleAmount[0] {
			dp[0][j] = j * wholeSalePrice[0]
		}
	}

	for i := 1; i < shopsCnt; i++ {
		for j := 0; j < cnt; j++ {
			for b := 0; b <= j; b++ {
				dp[i][j] = min(dp[i][j], getCost(retailPrice[i], wholeSalePrice[i], b, wholeSaleAmount[i])+dp[i-1][j-b], getCost(retailPrice[i], wholeSalePrice[i], j-b, wholeSaleAmount[i])+dp[i-1][b])
			}
		}
	}

	minVal := math.MaxInt
	for j := bricksCnt - 1; j < cnt; j++ {
		minVal = min(minVal, dp[shopsCnt-1][j])
	}

	return minVal
}

func getCost(retailPrice int, wholeSalePrice int, cnt int, wholeSaleAmount int) int {
	cost := retailPrice * cnt
	if cnt > wholeSaleAmount {
		cost = wholeSalePrice * cnt
	}
	return cost
}
