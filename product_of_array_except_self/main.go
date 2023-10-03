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
	cnt, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	multArr := make([]int, cnt)
	mult := 1

	for i := 0; i < cnt; i++ {
		multArr[i] = mult
		mult *= arr[i]
	}

	mult = 1
	for i := cnt - 1; i > -1; i-- {
		multArr[i] *= mult
		mult *= arr[i]
	}

	fmt.Println(multArr)

}
