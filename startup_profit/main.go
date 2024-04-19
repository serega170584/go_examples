package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n := scanner.Text()

	scanner.Scan()
	k := scanner.Text()

	scanner.Scan()
	d := scanner.Text()

	fmt.Println(getMaxStratupProfit(n, k, d))
}

func getMaxStratupProfit(n string, k string, d string) string {
	var sb strings.Builder

	nRunes := []rune(n)
	nIntLength := len(nRunes) + 1
	nInt := make([]int, nIntLength)

	for i := nIntLength - 1; i >= 1; i-- {
		nInt[i], _ = strconv.Atoi(string(nRunes[i-1]))
	}

	kRunes := []rune(k)
	kIntLength := len(kRunes)
	kInt := make([]int, kIntLength)

	dInt, _ := strconv.Atoi(d)

	for i, v := range kRunes {
		kInt[i], _ = strconv.Atoi(string(v))
	}

	if kIntLength == nIntLength {
		isEqual := true
		for i := 1; i < nIntLength; i++ {
			if kInt[i-1] != nInt[i] {
				isEqual = false
				break
			}
		}

		if isEqual == true {
			for i := 0; i < dInt-1; i++ {
				kInt = append(kInt, 0)
			}

			for _, v := range kInt {
				sb.WriteString(strconv.Itoa(v))
			}

			return sb.String()
		}
	}

	addInt := make([]int, nIntLength)
	copy(addInt, nInt)

	prevDigit := 0
	for i := 0; i < 9; i++ {
		for j := nIntLength - 1; j >= 0; j-- {
			sum := nInt[j] + addInt[j] + prevDigit
			nInt[j] = sum % 10
			prevDigit = sum / 10
		}
	}

	if len(k) > nIntLength {
		return "-1"
	}

	modNInt := make([]int, nIntLength)
	copy(modNInt, nInt)
	positiveModInt := make([]int, nIntLength)
	copy(positiveModInt, modNInt)

	prevBusyDigit := 0

	copyKVal := make([]int, nIntLength)
	for i, v := range kInt {
		copyKVal[i+nIntLength-kIntLength] = v
	}

	for i := nIntLength - 1; i >= 0; i-- {
		kVal := copyKVal[i]
		num := positiveModInt[i] - kVal - prevBusyDigit
		if num < 0 {
			num = positiveModInt[i] - kVal - prevBusyDigit + 10
			prevBusyDigit = 1
		} else {
			prevBusyDigit = 0
		}
		modNInt[i] = num
	}

	if prevBusyDigit == 1 {
		return "-1"
	}

	for prevBusyDigit != 1 {
		copy(positiveModInt, modNInt)
		for i := nIntLength - 1; i >= 0; i-- {
			kVal := copyKVal[i]
			num := modNInt[i] - kVal - prevBusyDigit
			if num < 0 {
				num = modNInt[i] - kVal - prevBusyDigit + 10
				prevBusyDigit = 1
			} else {
				prevBusyDigit = 0
			}
			modNInt[i] = num
		}
	}

	zeroCnt := 0
	for _, v := range positiveModInt {
		if v == 0 {
			zeroCnt++
		}
	}

	if zeroCnt != nIntLength {
		prevBusyDigit = 0
		for i := nIntLength - 1; i >= 0; i-- {
			kVal := copyKVal[i]
			num := kVal - positiveModInt[i] - prevBusyDigit
			if num < 0 {
				num = kVal - positiveModInt[i] - prevBusyDigit + 10
				prevBusyDigit = 1
			} else {
				prevBusyDigit = 0
			}
			positiveModInt[i] = num
		}

		for i, v := range positiveModInt {
			if i != nIntLength-1 && v != 0 {
				return "-1"
			}
		}
	}

	nInt[nIntLength-1] = nInt[nIntLength-1] + positiveModInt[nIntLength-1]

	for i := 0; i < dInt-1; i++ {
		nInt = append(nInt, 0)
	}

	for _, v := range nInt {
		sb.WriteString(strconv.Itoa(v))
	}

	return sb.String()
}
