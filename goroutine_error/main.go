package main

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

func main() {
	a := make([]int, 3)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(a []int, wg *sync.WaitGroup) {
		defer handler()
		defer wg.Done()
		//defer test(a)
		i := 3
		//defer testInt(i)
		//defer testIntPointer(&i)
		var u *int
		u = &i
		//defer testPointerUnnamed(*u)
		i = 2
		c := 9
		var b *int
		b = &c
		u = b
		fmt.Println(*u)
		fmt.Println(i)
		a[1] = 2
		//a[3] = 1
		a[2] = 2
	}(a, wg)
	wg.Wait()
	fmt.Println(a)
	var d int64
	var e int32
	intCheck(d, e)
	var f int64
	intCheck(d, f)

	fmt.Println(unsafe.Sizeof(a))

	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	fmt.Printf("stats: %d\n", stats.HeapAlloc)

	all := make([][]int, 0)

	//r := make([]int, 1<<24)
	r := mutate()
	fmt.Println(cap(r))

	runtime.GC()

	all = append(all, r)

	runtime.ReadMemStats(stats)
	fmt.Printf("stats after created: %d\n", stats.HeapAlloc)

	_ = r

}

func mutate() []int {
	r := make([]int, 1<<24)
	return r[1<<24-1:]
}

func test(a []int) {
	a[0] = 1
}

func testInt(i int) {
	fmt.Println(i)
}

func testIntPointer(i *int) {
	fmt.Println(*i)
}

func testPointerUnnamed(u int) {
	fmt.Println(u)
}

func handler() {
	e := recover()
	err, _ := e.(error)
	fmt.Println(err)
}

func intCheck(a interface{}, b interface{}) {
	fmt.Println(a == b)
}
