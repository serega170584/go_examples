package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type MinStorage struct {
	list    []int
	pointer int
}

func createMinStorage(cnt int) *MinStorage {
	m := &MinStorage{}
	m.pointer = cnt
	m.list = make([]int, cnt+1)
	m.list[cnt] = 1000000
	return m
}

func (m *MinStorage) add(num int) {
	m.pointer--
	m.list[m.pointer] = num
}

func (m *MinStorage) min() int {
	min := m.list[m.pointer]
	m.pointer++
	return min
}

func (m *MinStorage) sortArr() {
	sort.Ints(m.list)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	cntStr := scanner.Text()
	cnt, _ := strconv.Atoi(cntStr)

	arr := make([][]int, cnt)
	cnts := make([]int, cnt)
	pointers := make([]int, cnt)
	totalCnt := 0
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		curCnt, _ := strconv.Atoi(scanner.Text())

		cnts[i] = curCnt
		arr[i] = make([]int, curCnt+1)

		totalCnt += curCnt

		for j := 0; j < curCnt; j++ {
			scanner.Scan()
			arr[i][j], _ = strconv.Atoi(scanner.Text())
		}
		arr[i][curCnt] = 1000000
	}

	result := make([]int, totalCnt)
	resultPointer := 0

	m := createMinStorage(cnt)

	for i := 0; i < cnt; i++ {
		m.add(arr[i][0])
	}

	m.sortArr()

	min := m.min()

	var currentArrPointer int
	for min != 1000000 {
		for i := 0; i < cnt; i++ {
			if arr[i][pointers[i]] == min {
				pointers[i]++
				currentArrPointer = i
				result[resultPointer] = min
				resultPointer++
				break
			}
		}

		if pointers[currentArrPointer] != cnts[currentArrPointer] {
			el := arr[currentArrPointer][pointers[currentArrPointer]]
			m.add(el)
			m.sortArr()
		}

		min = m.min()
	}

	fmt.Println(result)
}
