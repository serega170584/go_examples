package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(convType(3))

	var i interface{} = "hello"
	t, ok := i.(float64)
	fmt.Println(t, ok)

	l := i.(float64)
	fmt.Println(l)

	r := "sadasdasdasd"
	z := []rune(r)
	y := string(z)
	fmt.Println(y)
	//for _, j := range y {
	//	fmt.Print(j)
	//}

	//in := i.(int) panic
	//fmt.Println(in)

	//fmt.Println(convStringToFloat64Interface(f))

	//x := make([]int, 4)
	//for i := 0; i < 4; i++ {
	//	x[i] = i + 1
	//}
	//fmt.Println(sliceMembers(x...)) panic

	//fmt.Println(convIntToFloat64Interface(64))

	//var b interface{} = "64"
	//c, ok := b.(float64)
	fmt.Println(convStringToFloat64Interface("64"))
	//b := 64
	//fmt.Println(c, ok)

	fmt.Println(int32To8(15))

	fmt.Printf("index data type:    %T\n", int32To8(15))

	fmt.Println(int64To8(64))
	fmt.Printf("index data type:    %T\n", int64To8(64))

	fmt.Println(int64To8(129))
	fmt.Printf("index data type:    %T\n", int64To8(129))

	fmt.Println(int64To8(130))
	fmt.Printf("index data type:    %T\n", int64To8(130))

	fmt.Println(int64To8(256))
	fmt.Printf("index data type:    %T\n", int64To8(256))

	fmt.Println(int64To8(383))
	fmt.Printf("index data type:    %T\n", int64To8(383))

	fmt.Println(int64To8(385))
	fmt.Printf("index data type:    %T\n", int64To8(385))

	fmt.Println(int64To8(384))
	fmt.Printf("index data type:    %T\n", int64To8(384))

	fmt.Println(1.79769313486231570814527423731704356798070e+308)
	a := 390.8
	fmt.Println(int(a))

	d := 1.79769313486231570814527423731704356798070e+10
	fmt.Println(int(d))
	fmt.Println(math.MaxInt64)

	e := "my string"

	f := []byte(e)

	g := string(f)

	fmt.Println(e)

	fmt.Println(f)

	fmt.Println(g)

}

func int32To8(i int32) int8 {
	return int8(i)
}

func int64To8(i int64) int8 {
	return int8(i)
}

func convType(i interface{}) (x int, ok bool) {
	x, ok = i.(int)

	return
}

//func convIntToFloat64(i int) (x float64) {
//	x = i.(float64)
//	return
//}
//
//func convStringToFloat64(i string) (x float64) {
//	x = i.(float64)
//	return
//}

func convIntToFloat64Interface(i interface{}) (x float64) {
	x = i.(float64)
	return
}

func convStringToFloat64Interface(i interface{}) (x float64, ok bool) {
	x, ok = i.(float64)
	return
}

//func convStringToFloat64Func(i interface{}) (x float64) {
//	x = float64(i)
//	return
//}

//func sliceMembers(x ...interface{}) (y int) {
//	for _, val := range x {
//		y = val.(int)
//	}
//	return
//}
