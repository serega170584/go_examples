package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	inPrefixList := strings.Split(scanner.Text(), " ")

	scanner.Scan()
	list := strings.Split(scanner.Text(), " ")

	prefixList := make(map[string]struct{}, len(inPrefixList))
	for _, v := range inPrefixList {
		prefixList[v] = struct{}{}
	}

	res := make([]string, 0, len(prefixList))
	for _, str := range list {
		cmpStr := ""
		for _, v := range str {
			cmpStr += string(v)
			if _, ok := prefixList[cmpStr]; ok {
				break
			}
		}
		res = append(res, cmpStr)
	}

	outRes := make([]any, 0, len(list))
	for _, v := range res {
		outRes = append(outRes, v)
	}
	fmt.Println(outRes...)
}
