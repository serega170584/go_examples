package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Params struct {
	x   []float64
	v   []float64
	n   int
	eps float64
}

func newParams(x []float64, v []float64, n int, eps float64) *Params {
	return &Params{x: x, v: v, n: n, eps: eps}
}

func (params *Params) X() []float64 {
	return params.x
}

func (params *Params) V() []float64 {
	return params.v
}

func (params *Params) N() int {
	return params.n
}

func (params *Params) Eps() float64 {
	return params.eps
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter list count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter epsilon")
	scanner.Scan()
	epsilon, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println("Enter speed list")
	v := make([]float64, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		v[i], _ = strconv.ParseFloat(scanner.Text(), 64)
	}

	fmt.Println("Enter distance list")
	x := make([]float64, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		x[i], _ = strconv.ParseFloat(scanner.Text(), 64)
	}

	maxV := v[0]
	minV := v[0]

	for i := 1; i < n; i++ {
		maxV = min(maxV, v[i])
		minV = min(minV, v[i])
	}

	minI := 0
	maxI := 0
	for i := 0; i < n; i++ {
		if v[i] == minV {
			minI = i
		}
		if v[i] == maxV {
			maxI = i
		}
	}

	r := 0.0000
	for t := 0; t < 1000000; t++ {
		minCurrentX := x[0] + v[0]*float64(t)
		maxCurrentX := x[0] + v[0]*float64(t)
		for i := 0; i < n; i++ {
			minCurrentX = min(x[i]+v[i]*float64(t), minCurrentX)
			maxCurrentX = max(x[i]+v[i]*float64(t), maxCurrentX)
		}

		if minCurrentX == x[minI]+v[minI]*float64(t) && maxCurrentX == x[maxI]+v[maxI]*float64(t) {
			r = float64(t)
		}
	}

	params := newParams(x, v, n, epsilon)

	fmt.Println("Got optimum time", lbsearch(0, r, epsilon, check, params))

}

func lbsearch(l float64, r float64, eps float64, check func(m float64, params *Params) bool, params *Params) float64 {
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
	return dist(m, params) <= dist(m+params.Eps(), params)
}

func dist(t float64, params *Params) float64 {
	x := params.X()
	v := params.V()
	minPos, maxPos := x[0]+v[0]*t, x[0]+v[0]*t
	for i := 1; i < params.N(); i++ {
		pos := x[1] + v[1]*t
		minPos = min(pos, minPos)
		maxPos = max(pos, maxPos)
	}
	return maxPos - minPos
}
