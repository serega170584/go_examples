package main

import (
	"fmt"
	"log"
)

// 5 35 40 101 59 63
// 6 35 40 101 59 63 101
// 4 35 40 101 0
// 7 35 40 101 59 63 101 40
// 7 35 40 101 59 53 101 40
// 8 35 40 101 59 53 101 70 80
// 8 35 40 101 90 53 101 70 80
// 1 35
// 8 35 101 90 101 70 101 80 90
// 8 35 101 0 0 0 0 0 0
func main() {
	var cnt int
	_, err := fmt.Scanln(&cnt)

	if err != nil {
		log.Fatal(err)
	}

	if cnt == 0 {
		fmt.Println(0)
		fmt.Println("0 0")
		return
	}

	dayPrices := make([]int, cnt+1)

	for i := 1; i < cnt+1; i++ {
		_, err = fmt.Scanln(&dayPrices[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	prices := make([][]int, cnt+1)
	for day := 0; day < cnt+1; day++ {
		prices[day] = make([]int, cnt+2)
		for couponInd := 0; couponInd < cnt+2; couponInd++ {
			prices[day][couponInd] = -1
			if day == 0 && couponInd == 0 {
				prices[day][couponInd] = 0
			}
		}
	}

	minCouponInd := 0
	minPrice := getPrice(cnt, 0, prices, dayPrices, cnt)
	for couponInd := 1; couponInd < cnt+2; couponInd++ {
		price := getPrice(cnt, couponInd, prices, dayPrices, cnt)
		if price < minPrice {
			minCouponInd = couponInd
			minPrice = price
		}
	}

	usedDays := getUsedDays(cnt, minCouponInd, prices, dayPrices, cnt)

	fmt.Println(prices)
	fmt.Println(minCouponInd)
	fmt.Println(usedDays)
}

func getPrice(day int, couponInd int, prices [][]int, dayPrices []int, cnt int) int {
	price := dayPrices[day]

	if prices[day][couponInd] != -1 {
		return prices[day][couponInd]
	}

	if day == 0 && day < couponInd {
		prices[day][couponInd] = 300
		return 300
	}

	if day == 0 {
		prices[day][couponInd] = 0
		return 0
	}

	if couponInd == cnt+1 {
		prices[day][couponInd] = getPrice(day-1, couponInd, prices, dayPrices, cnt) + 300
		return prices[day][couponInd]
	}

	if day < couponInd {
		prevDayPrice := getPrice(day-1, couponInd, prices, dayPrices, cnt) + 300
		minPrice := getPrice(day-1, couponInd+1, prices, dayPrices, cnt)
		if prevDayPrice < minPrice {
			minPrice = prevDayPrice
		}
		prices[day][couponInd] = minPrice
		return minPrice
	}

	if price > 100 && couponInd == 0 {
		prices[day][couponInd] = getPrice(day-1, couponInd+1, prices, dayPrices, cnt)
		return prices[day][couponInd]
	}

	minPrice := getPrice(day-1, couponInd+1, prices, dayPrices, cnt)
	prevDayPrice := getPrice(day-1, couponInd, prices, dayPrices, cnt) + price
	if price > 100 {
		prevDayPrice = getPrice(day-1, couponInd-1, prices, dayPrices, cnt) + price
	}

	if prevDayPrice < minPrice {
		minPrice = prevDayPrice
	}

	prices[day][couponInd] = minPrice
	return minPrice
}

func getUsedDays(day int, couponInd int, prices [][]int, dayPrices []int, cnt int) []int {
	if day == 0 || couponInd == cnt+1 {
		usedDays := make([]int, 0)
		return usedDays
	}

	price := dayPrices[day]
	prevCouponInd := couponInd
	if price > 100 {
		prevCouponInd = couponInd - 1
	}

	if prevCouponInd < 0 {
		return append(getUsedDays(day-1, couponInd+1, prices, dayPrices, cnt), day)
	}

	diff := prices[day][couponInd] - prices[day-1][prevCouponInd]
	if diff == price {
		return getUsedDays(day-1, prevCouponInd, prices, dayPrices, cnt)
	}

	return append(getUsedDays(day-1, couponInd+1, prices, dayPrices, cnt), day)
}
