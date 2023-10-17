package main

import "fmt"

func errorHandler(i int) {
	r := recover()
	err := r.(error)
	if err == nil {
		return
	}
	fmt.Println(err.Error(), " ", i)
}

func foo() {
	var i int
	i++
	defer errorHandler(i)
	i++
	smallSlice := []int{1, 0, 1}
	smallSlice[10] = 1
	i++
}

func main() {
	foo()
	fmt.Println("recovery, end of main")
}
