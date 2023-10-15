package main

import "fmt"

type Test struct {
	a int
	b string
}

func main() {
	test := getTest()
	if test == nil {
		fmt.Println("asdas")
	}
	checkTest(test)
}

func getTest() *Test {
	return nil
}

func checkTest(test *Test) {
	fmt.Println(test.a)
}
