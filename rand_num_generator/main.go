package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	g := randNumGenerator(cnt)
	for val := range g {
		fmt.Println(val)
	}
}

func randNumGenerator(n int) chan int {
	ch := make(chan int)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	go func() {
		for i := 0; i < n; i++ {
			ch <- r.Intn(n)
		}
		close(ch)
	}()

	return ch
}
