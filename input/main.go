package main

import (
	"fmt"
	"log"
)

func main() {
	var a, b int
	_, err := fmt.Scan(&a)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Scan(&b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a + b)
}
