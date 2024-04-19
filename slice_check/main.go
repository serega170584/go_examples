package main

import (
	"fmt"
	"runtime"
)

type A interface {
	test()
}

type TestA struct{}

func (a TestA) test() {}

func main() {
	var a A
	var b TestA
	fmt.Println(a == nil)
	a = b
	fmt.Println(a)
	//a := make([]int, 2, 2)
	//b := append(a, 1)
	//b[0] = 1
	//fmt.Println("New a", a)
	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	//fmt.Printf("stats: %d \n", stats.HeapAlloc)

	//a := [3]int{10, 20, 30}
	//fmt.Println(unsafe.Sizeof(a))

	runtime.ReadMemStats(stats)
	//fmt.Printf("stats: %d \n", stats.HeapAlloc)

	//b := []int{20}
	//fmt.Println(unsafe.Sizeof(b))
	//fmt.Println(reflect.ValueOf(b))
	//fmt.Println(reflect.TypeOf(b))
	//fmt.Println(reflect.ValueOf(b).Pointer(), *(*reflect.SliceHeader)(unsafe.Pointer(&b)))
	//
	//b = append(b, 30)
	//fmt.Println(unsafe.Sizeof(b))
	//fmt.Println(len(b))
	//
	//b = append(b, 50)
	//fmt.Println(unsafe.Sizeof(b))
	slice := make([]int, 2, 2)
	mutate(slice, 0)
	fmt.Println(slice)
}

func mutate(slice []int, i int) {
	slice[i] = 12
}
