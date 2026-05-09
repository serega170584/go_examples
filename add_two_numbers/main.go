package main

import "fmt"

func main() {
	a := []int{1, 2, 8, 9}
	b := []int{8, 9}

	l := max(len(a), len(b))

	a = append(make([]int, l-len(a)), a...)
	b = append(make([]int, l-len(b)), b...)

	res := make([]int, l)

	s := a[len(a)-1] + b[len(b)-1]
	extra := s / 10
	res[len(res)-1] = s % 10

	for k := len(res) - 2; k >= 0; k-- {
		s = a[k] + b[k] + extra
		extra = s / 10
		res[k] = s % 10
	}

	res[0] = extra

	fmt.Println(res)
}
