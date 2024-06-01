package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	image := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
	}

	search := [][]int{
		[]int{4, 5, 6},
		[]int{4, 5, 6},
		[]int{4, 5, 6},
	}

	fmt.Println(searchImage(8, 4, image, 3, 3, search))
}

func searchImage(iw int, ih int, image [][]int, sw int, sh int, search [][]int) (int, int) {
	m := make(map[string]bool)
	topRes := -1
	leftRes := -1
	for i, r := range image {
		for j, v := range r {
			for k, mv := range m {
				if mv {
					parts := strings.Split(k, "_")
					t, _ := strconv.Atoi(parts[0])
					l, _ := strconv.Atoi(parts[1])
					st := i - t
					sl := j - l
					if 0 <= st && st < sh && 0 <= sl && sl < sw && search[st][sl] != v {
						m[k] = false
						continue
					}
					if st == sh-1 && sl == sw-1 {
						m[k] = false
						topRes = i
						leftRes = j
					}
				}
			}
			if search[0][0] == v && i+sh-1 < ih && j+sw-1 < iw {
				k := strings.Join([]string{strconv.Itoa(i), strconv.Itoa(j)}, "_")
				m[k] = true
			}
		}
	}
	return topRes, leftRes
}
