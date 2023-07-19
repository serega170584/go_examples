package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var cnt int
	fmt.Println("Enter count")
	_, err := fmt.Scan(&cnt)
	if err != nil {
		log.Fatal(err)
	}

	x := make([]interface{}, cnt)

	fmt.Println("Enter array")

	a := make([][]string, cnt)
	for i := range a {
		a[i] = make([]string, cnt)
		for j := range a[i] {
			x[j] = &a[i][j]
		}
		_, err = fmt.Scanln(x...)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < cnt; i++ {
		for j := 0; j < i; j++ {
			a[i][j], a[j][i] = a[j][i], a[i][j]
		}
	}

	fmt.Println(a)

	middleCnt := cnt / 2
	fmt.Println(middleCnt)
	for i := 0; i < cnt; i++ {
		for j := 0; j < middleCnt; j++ {
			a[i][j], a[i][cnt-j-1] = a[i][cnt-j-1], a[i][j]
		}
	}

	for i := 0; i < cnt; i++ {
		y := make([]string, cnt)
		for j := 0; j < cnt; j++ {
			y[j] = fmt.Sprintf("%02s", a[i][j])
		}
		fmt.Println(strings.Join(y, " "))
	}
}
