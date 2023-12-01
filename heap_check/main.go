package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func main() {
	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	println("Memory is busied: ", stats.HeapAlloc)

	a := make([]int, 1<<25)

	_ = a

	//runtime.GC()

	runtime.ReadMemStats(stats)
	println("Memory is busied: ", stats.HeapAlloc)

	b := make([]int, 4)

	c := b[1:2]

	fmt.Println("B pointer: ", reflect.ValueOf(b).Pointer(), "C pointer: ", reflect.ValueOf(c).Pointer())
	fmt.Println("B unsafe pointer: ", unsafe.Pointer(&b), "C unsafe  pointer: ", unsafe.Pointer(&c))
	fmt.Println("B slice header: ", *(*reflect.SliceHeader)(unsafe.Pointer(&b)), "C slice header: ", *(*reflect.SliceHeader)(unsafe.Pointer(&c)))

	//a := make([]int, 1, 3)
	//b := a
	//b = append(b, 2)
	//b[0] = 1
	////append1(a, 1024)
	//fmt.Println(a)
	//fmt.Println(b)
}

//func append1(a []int, el int) {
//	b := a
//	a = append(a, el)
//	b[0] = 23
//	fmt.Println(b)
//}
