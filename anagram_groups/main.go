package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	arr := make([]string, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		arr[i] = scanner.Text()
	}

	anagramGroups := make(map[string][]string, cnt)
	for _, val := range arr {
		intBites := make([]int, len(val))
		for i, intBite := range val {
			intBites[i] = int(intBite)
		}
		sort.Ints(intBites)

		strBites := make([]string, len(val))
		for i, intBite := range intBites {
			strBites[i] = fmt.Sprintf("%c", intBite)
		}
		str := strings.Join(strBites, "")

		if _, ok := anagramGroups[str]; !ok {
			anagramGroups[str] = make([]string, 0)
		}
		anagramGroups[str] = append(anagramGroups[str], val)
	}

	fmt.Println(anagramGroups)

}
