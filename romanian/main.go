package main

import (
	"bufio"
	"fmt"
	"os"
)

// xix +
// xixi - we can't place i after ix
// xxi + 21
// xxiv + 24
// xxix + 29
// xxx + 30
// xxxix + 39
// xl + 40
// l + 50
// lx + 60
// lxx + 70
// lxxx + 80
// xc + 90
// c + 100
// d + 500
// m + 1000
// mmmcd + 3400
// mmmd + 3500
// mmmdc + 3600
// mmmdcc + 3700
// mmmdccc + 3800
// mmmcm + 3900
// mmmm - 4000
// mmmcmlxxxix  + 3989
// mm
// mmccxxii + 2222
// mmmi + 3001
// mmmcccxxxiii + 3333
// mmmcccxxxviii + 3338
// mmmdccclxxxviii + 3888
// mmmcccxxxiv + 3334
// mmcmlxxxix + 2989
// mcmlxix + 1969
// cmcd -
// cml +
// i v x l c d m
// i = 1 v = 5 x = 10 l = 50 c = 100 d = 500 m = 1000
// m d c l x v i
// m => c, d => c, c => x, l => x, x => i, v => i
// [m d c l x v i] = base - index pointer
// iv ix v mc vv viii xv xx xvi
// before cycle through base check if current symbol < accepted - so we validate number
// current sym begins with 5: check prev, if neighbor then fill current state, move to next state, move base pointer in 2 index break
// current sym begins with 1: check prev, if prev neighbour then fill current state, move to next state, move base pointer in 3 index break
// if current sym begins with 5 and and not empty prev sym in current state: move state
// if current sym begins with 5 move base pointer to next, increment count break
// fill state
// if count of symbol in current state = 3 then move current state to next, move base pointer to next break

// 13:17 - 14:11
// We should separate number
// number = state
// middle - start with 5
// border - start with 1
// 1. middle after previous border before - new state, for example, xv
// 2. middle after the same middle - not valid, for example, vv
// 3. border after previous middle before - the same state, for example, vi
// 4. border after the same border - the same state, for example, xx
// 5. middle after 1 next border after - fix state, move to next state, for example, iv
// 6. middle after >1 next border after - not valid, for example, iiv
// 7. border after 1 next border after - fix state, move to next state, for example, ix
// 8. border after >1 next border after - not valid, for example, iix
// 9. border after the same border before in previous state - not valid, for example, ixx
// 10. border after previous border before in previous state - not valid, for example, ixi
// 11. border after previous middle before in previous state - not valid, for example, ivi
// 12. border after the same border > 3 times - not valid, for example, xxxx
// 13. previous after next - not valid, lm
// 14. next after previous - next state, for example, mc

// We should interate number byte by byte
// Structures and responsibilities:
// 1. start - starting from which symbol we consider valid symbol for iteration, automatically drop not valid case before other checks of symbol(cases:2, 6, 8, 9, 10, 11, 12, 13)
// 2. states and counts matrix - counts symbols and helps define start symbol(1.)
// 4. dictionary of index - define index of current symbol in base ['m' => 0, 'd' => 1, ...]
// 5. corners - define starts form 1 or 5 symbol [0 => true, 1 => false,....]
// 6. given string of bytes
// 7. map romanian and arabian nums ['m' => 1000, 'd' => 500]
// 8. state number
// 9. sym index in given number
// 10. prev index
// 11. exclNum - exclused order rows - shows if state with non standard number or not

// cases we should write out conditions in iteration for: 1, 3, 4, 5, 7, 14

// Development
// 14:20 - 15:10
// 15:20 - 16:25

// Testing
// 16:25 - 16:33
// 18:18 - 18:56
// 18:58 - 19:49

func main() {

	scanner := makeScanner()
	scanner.Scan()
	romanianNum := scanner.Bytes()

	fmt.Println(getDefaultNum(romanianNum))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func makeArray(cnt int) [][]int {
	arr := make([][]int, cnt)
	for i := 0; i < cnt; i++ {
		arr[i] = make([]int, cnt)
	}
	return arr
}

// 6
func getDefaultNum(romNum []byte) int {
	// var 11
	exclNum := [7]bool{}

	// var 4
	symOrders := make(map[byte]int, 7)
	symOrders['M'] = 0
	symOrders['D'] = 1
	symOrders['C'] = 2
	symOrders['L'] = 3
	symOrders['X'] = 4
	symOrders['V'] = 5
	symOrders['I'] = 6

	// var 7
	base := [7]int{1000, 500, 100, 50, 10, 5, 1}

	// var 5
	corners := [7]bool{true, false, true, false, true, false, true}

	// var 2
	states := makeArray(7)

	// var 8
	stInd := 0
	// var 10
	prevInd := -1
	// var 1
	startInd := 0

	for _, romSym := range romNum {
		// var 9
		i := symOrders[romSym]

		if i < startInd {
			return -1
		}

		// case 7 9 10
		if corners[i] && i == prevInd-2 {
			states[stInd][i] = 1
			prevInd = i
			startInd = i + 3
			exclNum[stInd] = true
			stInd++
			continue
		}

		// case 5 2 11
		if !corners[i] && i == prevInd-1 {
			states[stInd][i] = 1
			prevInd = i
			startInd = i + 2
			exclNum[stInd] = true
			stInd++
			continue
		}

		// case 4 6 8
		if corners[i] && i == prevInd {
			states[stInd][i]++
			startInd = i
		}

		// case 12
		if states[stInd][i] == 3 {
			startInd = i + 1
			continue
		}

		// case 4 6 8
		if corners[i] && i == prevInd {
			continue
		}

		// case 3
		if corners[i] && i == prevInd+1 {
			states[stInd][i]++
			prevInd = i
			continue
		}

		if prevInd != -1 && states[stInd][prevInd] != 0 {
			stInd++
		}

		// case 1 14 13
		if corners[i] {
			states[stInd][i] = 1
			prevInd = i
			startInd = i - 2
			continue
		}

		// case 2
		states[stInd][i] = 1
		prevInd = i
		startInd = i + 1
	}

	var num int

	for stInd, cntRows := range states {
		rowSum := 0
		for i, cnt := range cntRows {
			curNum := cnt * base[i]
			if rowSum != 0 && exclNum[stInd] {
				num -= curNum
			} else {
				num += curNum
			}
			rowSum += curNum
		}
	}

	return num
}
