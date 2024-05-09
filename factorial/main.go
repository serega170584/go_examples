package main

import "fmt"

func main() {
	fmt.Println(factorial(10))
}

func factorial(n int) []int {
	res := make([]int, 0, 200)
	res = append(res, 1)
	if n == 1 {
		return res
	}

	s := make([]int, 1, 200)
	for i := 2; i <= n; i++ {
		m := i % 10
		d := i / 10
		extra := 0
		rl := len(res)
		for j := rl - 1; j >= 0; j-- {
			r := res[j]*m + extra
			s[j] = r % 10
			extra = r / 10
		}
		if extra > 0 {
			s = append([]int{extra}, s...)
		}

		sh := 1
		for d != 0 {
			m = d % 10
			d = d / 10
			extra = 0
			si := len(s) - 1 - sh
			ri := -1
			rl = len(res)
			for j := rl - 1; j >= 0; j-- {
				if si == -1 {
					ri = j
					break
				}
				r := res[j]*m + extra + s[si]
				s[si] = r % 10
				extra = r / 10
				si--
			}

			for j := ri; j >= 0; j-- {
				r := res[j]*m + extra
				s = append([]int{r % 10}, s...)
				extra = r / 10
			}

			if extra != 0 {
				s = append([]int{extra}, s...)
			}

			sh++
		}

		sl := len(s)
		res = res[:sl]
		copy(res, s)
	}

	return res
}
