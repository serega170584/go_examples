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
	cnt1, _ := strconv.Atoi(scanner.Text())
	arr1 := make([]int, cnt1)

	for i := 0; i < cnt1; i++ {
		scanner.Scan()
		arr1[i], _ = strconv.Atoi(scanner.Text())
	}

	scanner.Scan()
	cnt2, _ := strconv.Atoi(scanner.Text())
	arr2 := make([]int, cnt1+cnt2)

	for i := 0; i < cnt2; i++ {
		scanner.Scan()
		arr2[i], _ = strconv.Atoi(scanner.Text())
	}

	first := cnt1 - 1
	second := cnt2 - 1

	for i := cnt1 + cnt2 - 1; i > -1; i-- {
		if first == -1 {
			arr2[i] = arr2[second]
			second--
			continue
		}

		if second == -1 {
			arr2[i] = arr1[first]
			first--
			continue
		}

		if arr1[first] > arr2[second] {
			arr2[i] = arr1[first]
			first--
		} else {
			arr2[i] = arr2[second]
			second--
		}
	}

	fmt.Println(arr2)

}
