package main

import (
	"fmt"
	"log"
)

func main() {
	//var cnt int
	//fmt.Println("Enter count")
	//_, err := fmt.Scan(&cnt)
	//if err != nil {
	//	log.Fatal(err)
	//}
	x := make([]int, 3)
	y := make([]interface{}, 3)
	for i := range x {
		y[i] = &x[i]
	}
	_, err := fmt.Scanln(y...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(x)

	//formatString := make([]string, cnt)
	//for i := 0; i < cnt; i++ {
	//	formatString[i] = "%d"
	//}
	//format := strings.Join(formatString, " ")
	//
	//a := make([][]int, cnt)
	//for i := range a {
	//	a[i] = make([]int, cnt)
	//	for j := range a[i] {
	//		x := a[i]
	//		_, err := fmt.Scanf(format, x...)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//	}
	//}
	//fmt.Println("Input is done")
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Split(bufio.ScanWords)
	//var words []string
	//for scanner.Scan() {
	//	words = append(words, scanner.Text())
	//}
	//fmt.Println(words)
}
