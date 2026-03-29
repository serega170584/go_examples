package main

import (
	"fmt"
	"time"
)

0 1
1 2
2 3
3 4

func main() {
	t := time.Now()

	defer func() {
		fmt.Println(time.Since(t))
	}()

	tsl := make([]int, 1, 3)
	fmt.Println(tsl[:2])

	sl := make([]int, 1, 3)
	fmt.Println(sl) // [0] l=1, c=3

	appendSlice(sl, 1)
	fmt.Println(sl)     // [0] l=1, c=3
	fmt.Println(sl[:2]) // [0 1]

	copySlice(sl, []int{2, 3})
	fmt.Println(sl) // [0] l=1, c=3

	mutateSlice(sl, 1, 4) //error
	fmt.Println(sl)       // [2]
}

func appendSlice(sl []int, val int) {
	sl = append(sl, val)
}

func copySlice(sl, cp []int) {
	copy(sl, cp)
}

func mutateSlice(sl []int, ind, val int) {
	sl[ind] = val
}
