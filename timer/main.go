package main

import (
	"fmt"
	"sync"
)

func main() {
	//t := time.Now()
	//t1 := time.Now()
	//fmt.Println(t.After(t))
	m := make(map[struct{}]struct{})
	fmt.Printf("%p", m)
	fmt.Println()

	var mu sync.Mutex
	fmt.Printf("%p", &mu)
	fmt.Println()

	var wg sync.WaitGroup
	fmt.Printf("%p", &wg)
	fmt.Println()

	func(m1 map[struct{}]struct{}, mu1 sync.Locker, wg1 *sync.WaitGroup) {
		fmt.Printf("%p", m)
		fmt.Println()
		fmt.Printf("%p", m1)
		fmt.Println()
		fmt.Printf("%p", &mu)
		fmt.Println()
		fmt.Printf("%p", &mu1)
		fmt.Println()
		fmt.Printf("%p", &wg)
		fmt.Println()
		fmt.Printf("%p", &wg1)
	}(m, &mu, &wg)

}
