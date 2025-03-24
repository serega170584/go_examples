package main

import "fmt"

func main() {
	m := make(map[int]map[int]int)
	m[1][1] = 2
	fmt.Println(m)
}
