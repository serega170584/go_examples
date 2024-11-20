package main

import (
	"fmt"
	"sync"
)

type A interface {
	testFunc() A
}

type B struct{}

func (b B) testFunc() A {
	return b
}

func listen(name string, data map[string]string, c *sync.Cond, wg *sync.WaitGroup) {
	c.L.Lock()
	wg.Done()
	c.Wait()

	fmt.Printf("[%s] %s\n", name, data["key"])

	c.L.Unlock()
}

func broadcast(name string, data map[string]string, c *sync.Cond) {
	//time.Sleep(time.Second)

	c.L.Lock()

	data["key"] = "value"

	fmt.Printf("[%s] данные получены\n", name)

	c.Broadcast()
	c.L.Unlock()
}

type Test struct {
	test int
}

func (t *Test) testFunc(v int) {
	fmt.Printf("%p", t)
	t.test = v
	//debug.PrintStack()
	//debug.SetTraceback("10")
	//debug
}

func testFunc(v int) {
	t := Test{}
	t.testFunc(v)
}

func main() {
	m := make(map[int]int, 3)
	m[0] = 1
	m[1] = 2
	m[2] = 3
	m1 := make(map[int]int)
}
