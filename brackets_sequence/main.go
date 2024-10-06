package main

import (
	"bufio"
	"os"
	"strconv"
)

type Bracket struct {
	isOpen  bool
	isClose bool
	val     rune
}

type Stack struct {
	list []Bracket
}

func New(capacity int) *Stack {
	list := make([]Bracket, 0, capacity)
	return &Stack{list: list}
}

func (s *Stack) push(v Bracket) {
	l := len(s.list)
	s.list = s.list[:l+1]
	s.list[l] = v
}

func (s *Stack) pop() Bracket {
	l := len(s.list)
	v := s.list[l-1]
	s.list = s.list[:l-1]
	return v
}

func (s *Stack) empty() bool {
	return len(s.list) == 0
}

// ((()))
// (()())
// (())()
// ()(())
// ()()()
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	cntStr := scanner.Text()
	cnt, _ := strconv.Atoi(cntStr)

	result := generate(cnt)
	for _, row := range result {
		for _, v := range *row {
			print(string(v.val))
		}
		println()
	}
}

func generate(cnt int) []*[]Bracket {
	s := New(2 * cnt)
	list := make([]Bracket, 0, 2*cnt)

	for i := 0; i < cnt-1; i++ {
		s.push(Bracket{isOpen: true, val: []rune("(")[0]})
		list = append(list, Bracket{isOpen: true, val: []rune("(")[0]})
	}
	s.push(Bracket{isOpen: true, isClose: true, val: []rune("(")[0]})
	list = append(list, Bracket{isOpen: true, isClose: true, val: []rune("(")[0]})

	for i := 0; i < cnt; i++ {
		s.push(Bracket{isOpen: true, isClose: true, val: []rune(")")[0]})
		list = append(list, Bracket{isOpen: true, isClose: true, val: []rune(")")[0]})
	}

	result := make([]*[]Bracket, 0)
	listCopy := make([]Bracket, len(list))
	copy(listCopy, list)
	result = append(result, &listCopy)

	openCounter := cnt
	counter := 2 * cnt

	for !s.empty() {
		v := s.pop()
		list = list[:len(list)-1]

		if v.val == []rune("(")[0] {
			openCounter--
		}

		counter--

		if !v.isOpen {
			openCounter++
			if v.val == []rune("(")[0] {
				openCounter++
			}

			counter += 2

			isOpen := false
			if openCounter == cnt {
				isOpen = true
			}

			v.isOpen = true
			list = append(list, v)
			list = append(list, Bracket{isOpen: isOpen, isClose: false, val: []rune("(")[0]})
			s.push(v)
			s.push(Bracket{isOpen: isOpen, isClose: false, val: []rune("(")[0]})
		} else if !v.isClose {
			if v.val == []rune("(")[0] {
				openCounter++
			}

			counter += 2

			isOpen := false
			if openCounter == cnt {
				isOpen = true
			}

			isClose := false
			if openCounter == counter-openCounter {
				isClose = true
			}

			v.isClose = true
			list = append(list, v)
			list = append(list, Bracket{isOpen: isOpen, isClose: isClose, val: []rune(")")[0]})
			s.push(v)
			s.push(Bracket{isOpen: isOpen, isClose: isClose, val: []rune(")")[0]})
		}

		if len(list) == 2*cnt {
			func(listCopy []Bracket) {
				listCopy = make([]Bracket, len(list))
				copy(listCopy, list)
				result = append(result, &listCopy)
			}(listCopy)
		}
	}

	return result

}
