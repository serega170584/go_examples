package main

import "fmt"

func main() {
	a := []int{3, 2, 8, 9}
	b := []int{8, 9, 0, 0}

	l := max(len(a), len(b))
	a = append(make([]int, l-len(a)), a...)
	b = append(make([]int, l-len(b)), b...)

	res := make([]int, l+1)

	extra := 0
	for j := l - 1; j >= 0; j-- {
		v := a[j] + b[j] + extra
		res[j+1] = v % 10
		extra = v / 10
	}

	res[0] = extra
	if res[0] == 0 {
		res = res[1:]
	}

	fmt.Println(res)
}
