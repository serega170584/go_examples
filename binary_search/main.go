package main

func main() {
	arr := make([]int, 3)
	left := 0
	right := len(arr)
	for left+1 < right {
		mid := (left + right) / 2
		if arr[mid] <= 0 {
			left = mid
		} else {
			right = mid
		}
	}
}
