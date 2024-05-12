package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"", "ab", "abcde"}))
}

func longestCommonPrefix(a []string) string {
	al := len(a)
	l := len(a[0])
	for i := 1; i < al; i++ {
		l = min(l, len(a[i]))
	}
	prefix := make([]int32, 0, l)

	r := make([][]int32, al)
	for i := 0; i < al; i++ {
		r[i] = make([]int32, l)
		r[i] = []int32(a[i][:l])
	}

	for i := 0; i < l; i++ {
		v := r[0][i]
		isNotEqual := false
		for j := 1; j < al; j++ {
			if r[j][i] != v {
				isNotEqual = true
				break
			}
		}
		if isNotEqual {
			break
		}
		prefix = append(prefix, v)
	}

	return string(prefix)
}
