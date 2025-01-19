package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{[]int{12}}))
}

func spiralOrder(A [][]int) []int {
	res := make([]int, 0, len(A[0])*len(A))
	if len(A) == 1 {
		return A[0]
	}
	if len(A[0]) == 1 {
		for i := 0; i < len(A); i++ {
			res = append(res, A[i][0])
		}
		return res
	}
	maxColNum := len(A[0]) - 1
	maxRowNum := len(A) - 1
	maxRowLayer := (len(A) - 1) / 2
	maxColLayer := (len(A[0]) - 1) / 2
	maxLayer := maxRowLayer
	if maxColLayer < maxLayer {
		maxLayer = maxColLayer
	}

	for layer := 0; layer <= maxLayer; layer++ {
		for i := layer; i <= maxColNum-layer; i++ {
			res = append(res, A[layer][i])
		}
		if layer+1 > maxRowNum-layer {
			break
		}
		for i := layer + 1; i <= maxRowNum-layer; i++ {
			res = append(res, A[i][maxColNum-layer])
		}
		if layer > maxColNum-layer-1 {
			break
		}
		for i := maxColNum - layer - 1; i >= layer; i-- {
			res = append(res, A[maxRowNum-layer][i])
		}
		for i := maxRowNum - layer - 1; i > layer; i-- {
			res = append(res, A[i][layer])
		}
	}

	return res
}
