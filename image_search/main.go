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
	searched := make(map[string]bool, len(image))
	topRes := -1
	leftRes := -1
	for i, row := range image {
		for j, v := range row {
			for sk, s := range searched {
				if s == true {
					parts := strings.Split(sk, "_")
					top, _ := strconv.Atoi(parts[0])
					left, _ := strconv.Atoi(parts[1])
					tsp := i - top
					lsp := j - left
					if tsp == sh-1 && lsp == sw-1 && search[tsp][lsp] == v {
						topRes = i
						leftRes = j
						searched[sk] = false
						continue
					}
					if 0 <= tsp && tsp < sh && 0 <= lsp && lsp < sw && search[tsp][lsp] != v {
						searched[sk] = false
						continue
					}
					if tsp > sh-1 {
						continue
					}
					if lsp > sw-1 {
						continue
					}
				}
			}
			if v == search[0][0] {
				if i+sh-1 < ih && j+sw-1 < iw {
					searched[strings.Join([]string{strconv.Itoa(i), strconv.Itoa(j)}, "_")] = true
				}
			}
		}
	}

	return topRes, leftRes
}
