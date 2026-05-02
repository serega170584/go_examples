package main

import "fmt"

//		 99
//		 99
//		 81
//	 81
//	 891
//	 81
//
// 1701
// 99 * 99 = (90 + 9) *(90 + 9) = 90 * 90 + 90 * 9 + 9 * 90 + 9 * 9
func main() {
	a := []int{9, 9}
	b := []int{9, 9}
	c := make([]int, len(a)+len(b))
	for i := len(a) - 1; i >= 0; i-- {
		for j := len(b) - 1; j >= 0; j-- {
			res := c[i+j+1] + a[i]*b[j]
			lastRes := res % 10
			firstRes := res / 10
			c[i+j+1] = lastRes
			c[i+j] += firstRes
		}
	}
	fmt.Println(c)
}
