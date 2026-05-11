package main

import "fmt"

func main() {
	a := []int{9, 9}
	b := []int{9, 9}
	c := make([]int, len(a)+len(b))
	for i := len(a) - 1; i >= 0; i-- {
		for j := len(b) - 1; j >= 0; j-- {
			v := a[i]*b[j] + c[i+j+1]
			c[i+j+1] = v % 10
			c[i+j] += v / 10
		}
	}

	fmt.Println(c)
}
