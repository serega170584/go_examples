package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Params struct {
	percent  float64
	sum      float64
	mPercent float64
}

func NewParams(percent float64, sum float64) *Params {
	return &Params{percent: percent, sum: sum}
}

func (params *Params) Percent() float64 {
	return params.percent
}

func (params *Params) setMPercent(percent float64) {
	params.mPercent = percent
}

func (params *Params) getMPercent() float64 {
	return params.mPercent
}

func (params *Params) Sum() float64 {
	return params.sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter year percent")
	scanner.Scan()
	yearPercent, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println("Credit sum")
	scanner.Scan()
	creditSum, _ := strconv.ParseFloat(scanner.Text(), 64)

	eps := 0.0001

	params := NewParams(yearPercent, creditSum)
	mPercent := binarySearch(0, yearPercent, eps, check, params)
	fmt.Println("Got month percent", mPercent)

	params.setMPercent(mPercent)
	fmt.Println("Got month pay", binarySearch(0, creditSum, eps, checkPaySum, params))
}

func binarySearch(l float64, r float64, eps float64, check func(m float64, params *Params) bool, params *Params) float64 {
	for l+eps < r {
		m := (l + r) / 2
		if check(m, params) {
			r = m
		} else {
			l = m
		}
	}
	return l
}

func check(m float64, params *Params) bool {
	msum := 1 + m/100
	ysum := 1 + params.Percent()/100
	return math.Pow(msum, 12) >= ysum
}

func checkPaySum(m float64, params *Params) bool {
	periods := 36
	creditSum := params.Sum()
	mpercent := params.getMPercent()
	for i := 0; i < periods; i++ {
		mSum := creditSum * mpercent / 100
		creditSum -= m - mSum
	}
	return creditSum <= 0
}
