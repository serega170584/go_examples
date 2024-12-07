package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

const (
	empty = iota
	sign
	openBracket
	closedBracket
	number
)

type Stack struct {
	list []string
}

func (st *Stack) push(v string) {
	st.list = append(st.list, v)
}

func (st *Stack) pop() *string {
	if len(st.list) == 0 {
		return nil
	}

	item := st.list[len(st.list)-1]
	st.list = st.list[:len(st.list)-1]
	return &item
}

func NewStack(capacity int) *Stack {
	return &Stack{list: make([]string, 0, capacity)}
}

// 123 456
func main() {
	fmt.Println(strconv.Atoi("123 456"))
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := scanner.Text()
	n := 0
	numberSt := NewStack(utf8.RuneCountInString(t))
	signStack := NewStack(utf8.RuneCountInString(t))

	nm := map[rune]struct{}{'0': {}, '1': {}, '2': {}, '3': {}, '4': {}, '5': {}, '6': {}, '7': {}, '8': {}, '9': {}}

	priority := map[rune]int{'+': 0, '-': 0, '*': 1}

	prevType := empty
	token := ""
	for _, v := range t {
		if v == '+' || v == '-' || v == '*' {
			if prevType == number {
				signItem := signStack.pop()
				if signItem != nil {
					if priority[v] > priority[[]rune(*signItem)[0]] {
						signStack.push(*signItem)
						signStack.push(string(v))
						numberSt.push(token)
						token = string(v)
						prevType = sign
					} else {

					}
				}
			}
			fmt.Println("WRONG")
			return
		}
	}
}
