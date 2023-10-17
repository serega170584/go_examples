package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//func add(arr []int, v int) {
//	arr = append(arr, v)
//}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for {
		scanner.Scan()

		cnt, _ := strconv.Atoi(scanner.Text())
		for i := 0; i < cnt; i++ {
			err := worker(i)
			if err != nil {
				fmt.Printf("%s", err.Error())
			}
		}
	}
}

func worker(i int) error {
	mod := (i + 10) % 10
	if mod == 0 {
		return fmt.Errorf(fmt.Sprintf("Error %d\n", i+10))
	} else {
		fmt.Println(i + 10)
	}
	return nil
}
