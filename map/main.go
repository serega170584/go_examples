package main

import "fmt"

func main() {
	defer func() {
		e := recover()
		if e != nil {
			_, ok := e.(error)
			if ok {
				fmt.Println("error")
			}
		}
	}()

	m := make(map[int]int)

	m[1] = 1

	mutateMap(m)

	fmt.Println(m)

	sl := make([]int, 0, 1)

	mutateSlice(sl)

	fmt.Println(sl)

	sl[0] = 1
}

func mutateMap(m map[int]int) {
	m[2] = 2
}

func mutateSlice(sl []int) {
	sl = append(sl, 2)
}
