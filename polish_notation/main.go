package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stack struct {
	capacity int
	length   int
	list     []string
}

func NewStack(n int) *Stack {
	list := make([]string, n)
	return &Stack{capacity: n, list: list}
}

func (s *Stack) push(el string) {
	ind := s.length
	s.list[ind] = el
	s.length++
}

func (s *Stack) pop() string {
	length := s.length
	ind := length - 1
	s.length = ind
	return s.list[ind]
}

func getAnswer(n int, expr []string) string {
	polish, length := getPolish(n, expr)
	s := NewStack(n)
	for i := 0; i < length; i++ {
		v := polish[i]
		if v != "*" && v != "/" && v != "+" && v != "-" {
			s.push(v)
		} else {
			x := s.pop()
			y := s.pop()
			res := calcExpression(v, x, y)
			s.push(res)
		}
	}
	return s.pop()
}

func calcExpression(v string, x string, y string) string {
	xVal, _ := strconv.Atoi(x)
	yVal, _ := strconv.Atoi(y)

	if v == "*" {
		return strconv.Itoa(xVal * yVal)
	}

	if v == "/" {
		return strconv.Itoa(xVal + yVal)
	}

	if v == "+" {
		return strconv.Itoa(xVal + yVal)
	}

	return strconv.Itoa(xVal - yVal)
}

func getPolish(n int, expr []string) ([]string, int) {
	s := NewStack(n)
	polish := make([]string, n)
	polishInd := 0
	priorities := getPriorities()
	for _, v := range expr {
		if v == "*" || v == "/" || v == "+" || v == "-" {
			if s.length == 0 {
				s.push(v)
			} else {
				lastSym := s.pop()
				if priorities[v] <= priorities[lastSym] {
					polish[polishInd] = lastSym
					polishInd++
				} else {
					s.push(lastSym)
				}
				s.push(v)
			}
			continue
		}

		if v == "(" {
			s.push(v)
		}

		if v == ")" {
			lastSym := s.pop()
			for lastSym != "(" {
				polish[polishInd] = lastSym
				polishInd++
				lastSym = s.pop()
			}
		}

		if v == "1" || v == "2" || v == "3" || v == "4" || v == "5" || v == "6" || v == "7" || v == "8" || v == "9" {
			polish[polishInd] = v
			polishInd++
		}
	}

	for s.length > 0 {
		polish[polishInd] = s.pop()
		polishInd++
	}

	return polish, polishInd
}

func getPriorities() map[string]int {
	m := make(map[string]int, 4)
	m["*"] = 2
	m["/"] = 2
	m["+"] = 1
	m["-"] = 1
	return m
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter text length")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter expression")
	scanner.Scan()
	expr := scanner.Text()
	stringList := make([]string, n)
	for i, v := range expr {
		stringList[i] = string(v)
	}

	fmt.Println(getAnswer(n, stringList))
}
