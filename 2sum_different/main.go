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

	fmt.Println("Enter sum")
	scanner.Scan()
	sum, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list")
	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	num, num1 := twoSum(sum, n, list)

	fmt.Println("Sum numbers ", num, " ", num1)

}

func twoSum(sum int, n int, list []int) (int, int) {
	existed := make(map[int]struct{}, n)
	for _, v := range list {
		num := sum - v
		if _, ok := existed[num]; ok {
			return v, num
		}
		existed[v] = struct{}{}
	}
	return 0, 0
}
