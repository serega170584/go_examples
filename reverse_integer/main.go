package main

import (
	"fmt"
	"math"
)

func main() {
	var x int32 = -120
	var res int32 = 0
	for x != 0 {
		digit := x % 10
		x /= 10

		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && digit > 7) {
			fmt.Println(0)
		}

		if res < math.MinInt32/10 || (res == math.MinInt32/10 && digit < 8) {
			fmt.Println(0)
		}

		res = res*10 + digit
	}

	fmt.Println(res)
}
