package main

import (
	"fmt"
	"time"
)

type data struct {
	Body  string
	Error error
}

func doGet(url string) (string, error) {
	time.Sleep(time.Millisecond * 200)
	return fmt.Sprintf("Response of %s", url), nil
}

func future(url string) <-chan data {
	c := make(chan data, 1)

	go func() {
		body, err := doGet(url)

		c <- data{Body: body, Error: err}
	}()

	return c
}

func main() {
	future1 := future("http://example1.com")
	future2 := future("http://example2.com")

	fmt.Println("Request started")

	body1 := <-future1
	body2 := <-future2

	fmt.Printf("Response 1: %v\n", body1)
	fmt.Printf("Response 2: %v\n", body2)
}
