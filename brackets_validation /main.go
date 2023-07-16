package main

import "unicode/utf8"

func main() {

}

func isValid(s string) bool {
	stack := make([]string, utf8.RuneCountInString(s))
	currentIndex := 0
	for _, val := range s {
		symbol := string(val)
		if symbol == "{" || symbol == "(" || symbol == "[" {
			stack[currentIndex] = symbol
			currentIndex++
		}

		if (symbol == "}" || symbol == ")" || symbol == "]") && currentIndex == 0 {
			return false
		}

		if symbol == "}" {
			if stack[currentIndex-1] == "{" {
				currentIndex--
			} else {
				return false
			}
		} else if symbol == ")" {
			if stack[currentIndex-1] == "(" {
				currentIndex--
			} else {
				return false
			}
		} else if symbol == "]" {
			if stack[currentIndex-1] == "[" {
				currentIndex--
			} else {
				return false
			}
		}
	}
	return currentIndex == 0
}
