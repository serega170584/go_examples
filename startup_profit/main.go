package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	d, _ := strconv.Atoi(scanner.Text())

	add := strconv.Itoa(n)
	if n == k {
		add += "0"
		d--
	} else {
		r := n
		for i := 0; i < 9; i++ {
			r += n
			r = r % k
		}
		if r == 0 {
			add += "0"
		} else {
			addNum := k - r
			if addNum > 9 {
				fmt.Println("-1")
				return
			}
			add += strconv.Itoa(addNum)
		}
		d--
	}

	for d > 0 {
		add += "0"
		d--
	}

	fmt.Println(add)
}
