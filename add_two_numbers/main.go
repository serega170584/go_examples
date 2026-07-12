package main

import "fmt"

func main() {
	a := []int{3, 2, 8, 9}
	b := []int{8, 9, 0, 0}

	l := max(len(a), len(b))
	a = append(make([]int, l-len(a)), a...)
	b = append(make([]int, l-len(b)), b...)

	res := make([]int, l)
	extra := 0
	for i := l - 1; i >= 0; i-- {
		v := a[i] + b[i] + extra
		res[i] = v % 10
		extra = v / 10
	}

	if extra != 0 {
		res = append([]int{extra}, res...)
	}

	fmt.Println(res)
}
