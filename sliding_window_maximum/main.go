package main

import "fmt"

func main() {
	a := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	res := make([]int, 0)
	indexes := make([]int, k)
	indexes[0] = 0
	for i := 1; i < k; i++ {
		l := len(indexes)
		for j := 0; j < l; j++ {
			index := indexes[j]
			if a[i] >= a[index] {
				indexes = indexes[:j]
				break
			}
		}
		indexes = append(indexes, i)
	}

	res = append(res, a[indexes[0]])
	for i := k; i < len(a); i++ {
		index := indexes[0]
		if i-index == k {
			indexes = indexes[1:]
		}

		l := len(indexes)
		for j := 0; j < l; j++ {
			index = indexes[j]
			if a[i] >= a[index] {
				indexes = indexes[:j]
				break
			}
		}
		indexes = append(indexes, i)
		res = append(res, a[indexes[0]])
	}

	fmt.Println(res)
}
