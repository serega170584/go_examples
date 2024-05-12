package main

import "fmt"

func main() {
	s1 := "abc#"
	s2 := "ab"
	fmt.Println(isEqual(s1, s2))
}

func isEqual(s1 string, s2 string) bool {
	is1 := []int32(s1)
	is2 := []int32(s2)
	sl1 := len(is1)
	sl2 := len(is2)
	sl := max(sl1, sl2)
	sc1 := 0
	sc2 := 0
	ns1 := make([]int32, 0, sl1)
	ns2 := make([]int32, 0, sl2)
	nsp1 := 0
	nsp2 := 0
	for i := sl - 1; i >= 0; i-- {
		if i < sl1 {
			v := is1[i]
			if v == []int32("#")[0] {
				sc1++
			} else {
				if sc1 == 0 {
					ns1 = append(ns1, v)
				} else {
					sc1--
				}
			}
		}

		if i < sl2 {
			v := is2[i]
			if v == []int32("#")[0] {
				sc2++
			} else {
				if sc2 == 0 {
					ns2 = append(ns2, v)
				} else {
					sc2--
				}
			}
		}

		if nsp1 <= len(ns1)-1 && nsp2 <= len(ns2)-1 {
			if ns1[nsp1] != ns2[nsp2] {
				return false
			}
			nsp1++
			nsp2++
		}
	}

	if len(ns1) != len(ns2) {
		return false
	}

	return true
}
