package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	a := scanner.Text()

	scanner.Scan()
	b := scanner.Text()

	r := isSameAnagram(a, b)

	if r {
		fmt.Println("YES")
		return
	}

	fmt.Println("NO")

}

func isSameAnagram(a string, b string) bool {
	ar := []rune(a)
	br := []rune(b)

	if len(ar) != len(br) {
		return false
	}

	slices.Sort(ar)
	slices.Sort(br)

	for i := range ar {
		if ar[i] != br[i] {
			return false
		}
	}

	return true
}
