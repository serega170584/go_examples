package main

import (
	"fmt"
)

type Err1 struct{}

func (e *Err1) Error() string {
	return "123"
}

func err() error {
	var err1 error
	return err1
}

func err1() *error {
	var err1 *error
	return err1
}

//func err2() error {
//	var err1 Err1
//	return err1
//}

func err3() error {
	var err1 *Err1
	return err1
}

type Example struct {
	a int
	b int
}

func (e Example) setA(a int) {
	e.a = a
}

func (e *Example) setB(b int) {
	e.b = b
}

func changePointer(a *Example) {
	//a = nil
	*a = Example{a: 1, b: 2}
}

func add(sl []int) {
	sl = append(sl, 1)
	sl = append(sl, 2)
	sl = append(sl, 3)
}

func change(sl []int) {
	sl[0] = 1
	sl[1] = 2
	sl = append(sl, 3)
	sl[0] = 4
}

func deferFunc(a int) {
	fmt.Println(a)
}

func getSubSlice(sl []int) []int {
	last := len(sl) - 1
	//return sl[last:]
	v := make([]int, 1)
	copy(v, sl[last:])
	return v
}

func main() {
	//var e Example
	//e.setA(1)
	//fmt.Println(e.a)

	//var e1 Example
	//e1.setB(2)
	//fmt.Println(e1.b)

	//var a interface{}
	//fmt.Println(a == nil)
	//
	//var b *int
	//a = b
	//fmt.Println(a == nil)

	//fmt.Println(a.(*int) == nil)

	//fmt.Println(err() == nil)
	//fmt.Println(err1() == nil)
	//fmt.Println(err2() == nil)
	//fmt.Println(err3() == nil)

	//a := "qwerty字"
	//for _, v := range a {
	//	fmt.Println(string(v))
	//}
	//
	//a += "asdas字 字 字 字 "
	//
	//var sb strings.Builder
	//sb.WriteString(a)
	//sb.WriteString("asdas字 字 字 字 ")
	//
	//fmt.Println(sb.String())

	//a := &Example{}
	//changePointer(a)
	//fmt.Println(a)

	//sl := make([]int, 0, 2)
	//
	//add(sl)
	//
	//fmt.Println(sl)

	sl := make([]int, 2)

	change(sl)

	fmt.Println(sl)

	sl := make([]int, 1000)
	getSubSlice(sl)

	//a := 123
	//defer func() {
	//	fmt.Println(a)
	//}()

	//defer deferFunc(a)
	//a = 456

	//var b *int
	//b = &a
	//defer func() {
	//	fmt.Println(*b)
	//}()
	//*b = 456

	//counter := 0
	//var counter int64 = 0
	//wg := &sync.WaitGroup{}
	//wg.Add(1000)
	//for i := 0; i < 1000; i++ {
	//	go func() {
	//defer wg.Done()
	//counter++
	//atomic.AddInt64(&counter, 1)
	//}()
	//}
	//wg.Wait()

	//fmt.Println(counter)
}
