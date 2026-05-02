package main

import "fmt"

func main() {
	image := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	n := 4
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - layer - 1
		for i := first; i < last; i++ {
			offset := i - first

			tmp := image[first][i]

			image[first][i] = image[last-offset][first]
			image[last-offset][first] = image[last][last-offset]
			image[last][last-offset] = image[i][last]
			image[i][last] = tmp
		}
	}

	fmt.Println(image)
}
