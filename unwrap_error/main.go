package main

import (
	"errors"
	"fmt"
)

var (
	errUhOh = errors.New("oh critical error!!")
)

func check(num int) error {
	if num == 1 {
		return fmt.Errorf("it's odd")
	} else if num == 2 {
		return errUhOh
	}
	return nil
}

func validations(num int) error {

	err := check(num)
	if err != nil {
		return fmt.Errorf("run error: %w", err)
	}
	return nil
}

func main() {
	for num := 1; num <= 5; num++ {
		fmt.Printf("validating %d... ", num)
		err := validations(num)
		if err == errUhOh || errors.Unwrap(err) == errUhOh {
			fmt.Println("oh no something has happened!")
		} else if err != nil {
			fmt.Println("some error is present...", err)
		} else {
			fmt.Println("valid number only...!")
		}
	}
}
