package main

import "fmt"

func main() {
	a := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	mi := 0
	for i := 1; i < k; i++ {
		if a[i] > a[mi] {
			mi = i
		}
	}

	maxIndexes := make([]int, 0, len(a)-k)
	maxIndexes = append(maxIndexes, mi)

	res := make([]int, 0, len(a)-k)
	res = append(res, a[mi])

	for i := k; i < len(a); i++ {
		if i-maxIndexes[0] == k {
			maxIndexes = maxIndexes[1:]
		}

		for j := len(maxIndexes) - 1; j >= 0; j-- {
			index := maxIndexes[j]
			if a[i] < a[index] {
				break
			}

			maxIndexes = maxIndexes[:len(maxIndexes)-1]
		}

		maxIndexes = append(maxIndexes, i)

		index := maxIndexes[0]
		res = append(res, a[index])
	}

	fmt.Println(res)
}
