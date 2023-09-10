package main

import (
	"fmt"
	"log"
)

// 07:08 - 08:01
// abBa - aa
// aBba - aa
// aBbA - empty
// ABba - empty
// Abba - Abba
// ABbaA - A
// ABbaB - B
// ABbaCc - empty
// if prev == -1 prev++ next++
// create array filled by bool with len = len of string - array of exclused symbols
// take prev pointer, next pointer
// compare next pointer and prev pointer
// if next - prev in absolute = 32 then next and prev set true in array of exclused symbols decrement prev increment next
// 1. prev == -1 prev++ next++
// 2. next - prev == 32
// 3. next - prev > 1 prev = next next++
// variables:
// 1. string
// 2. prev
// 3. next
// 4. exclusedSyms
// 08:16 - 08:43
func main() {
	//1
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		log.Fatal(err)
	}
	//2
	prev := -1
	//3
	//next := 0
	//4
	exclusedSyms := make([]bool, len(str))
	for next := range str {
		if prev == -1 {
			prev = next
			continue
		}

		b := str[next]
		a := str[prev]
		diff := b - a
		if a > b {
			diff = a - b
		}
		if diff == 32 {
			exclusedSyms[prev] = true
			exclusedSyms[next] = true
			prev--
			continue
		}

		prev = next
	}

	for i, sym := range str {
		if !exclusedSyms[i] {
			fmt.Printf("%c", sym)
		}
	}
}
