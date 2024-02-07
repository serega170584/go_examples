package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type HashTable struct {
	list        [][]int
	lastIndexes []int
	cnt         int
}

func makeHashTable(cnt int) *HashTable {
	list := make([][]int, cnt)
	lastIndexes := make([]int, cnt)
	for i := range list {
		list[i] = make([]int, cnt)
	}
	return &HashTable{list: list, cnt: cnt, lastIndexes: lastIndexes}
}

func (ht *HashTable) Add(val int) {
	hash := val % ht.cnt
	lastIndex := ht.lastIndexes[hash]
	for i := 0; i < lastIndex; i++ {
		if ht.list[hash][i] == val {
			return
		}
	}
	ht.list[hash][lastIndex] = val
	lastIndex++
	ht.lastIndexes[hash] = lastIndex
}

func (ht *HashTable) Find(val int) bool {
	hash := val % ht.cnt
	lastIndex := ht.lastIndexes[hash]
	for i := 0; i < lastIndex; i++ {
		if val == ht.list[hash][i] {
			return true
		}
	}
	return false
}

func (ht *HashTable) Delete(val int) {
	hash := val % ht.cnt
	lastIndex := ht.lastIndexes[hash]
	for i := 0; i < lastIndex; i++ {
		if val == ht.list[hash][i] {
			if lastIndex > 1 {
				ht.list[hash][i] = ht.list[hash][lastIndex-1]
			}
			ht.lastIndexes[hash]--
			return
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list")
	ht := makeHashTable(n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		ht.Add(val)
	}
	fmt.Println(ht)

	fmt.Println("Add number")
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	ht.Add(num)
	fmt.Println(ht)

	fmt.Println("Find number")
	scanner.Scan()
	num, _ = strconv.Atoi(scanner.Text())
	fmt.Println(ht.Find(num))

	fmt.Println("Delete number")
	scanner.Scan()
	num, _ = strconv.Atoi(scanner.Text())
	ht.Delete(num)
	fmt.Println(ht)

	fmt.Println("Add number")
	scanner.Scan()
	num, _ = strconv.Atoi(scanner.Text())
	ht.Add(num)
	fmt.Println(ht)
}
