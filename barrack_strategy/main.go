package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 2500 5000 2499
// 2500 2500 0
// 2500 2500 2499
// 2500 2000 499
// 2001 2000 499
// 2001 2000 2998

// 2500 2300 299
// 2201 2300 299

// 2500 1500 999
// 1501 1500 999

// 2500 2400 99
// 2401 2200 2598

// 2500 2490 9
// 2491 2490 9

// 2500 2499 0
// 2500 2499 0

// 2500 2500 2499

// 2500 2499 2499

// 2500 2499 2499

// 2496 2495 2503

// 2401 2400 2598

// 2551 2450 2548

// 10 21 15
// 10 16 10

// 0 ... 10
// 0 ... 5 ... 10

// 10 18 8

// 100 110 150

// 10 1 15

// 3 4 3
// 1 1 1
// 1 0 1
//

// 1 1 1

// 3 2 2
// 3 0 2
// 3 0 0

// 4 3 5
// 4 0 5
// 4 0 1
// 3 0 1
// 2 0 0

// 4 4 5
// 4 0 5

// 4 5 5
// 4 1 5
// 3 0 7

// 4 5 5
// 4 1 5
// 4 1 6
// 4 1 2
// 2 1 2
// 2 1 7
// 1 0 1

// 4 8 5
// 4 4 5
// 4 3 2
// 2 3 2
// 2 3 5

// 4 4 7
// 4 0 7
// 4 0 3
// 1 0 10

// 4 1 6
// 4 0 3
// 1 0 3

// 4 1 6
// 4 1 2
// 2 1 2

// 4 3 4
// 4 3 0

// 4 3 4
//

// 2500 2500 2499

// 2500 2498 1
// 2499 2498 2500

// 2500 2500 2499
// 2498 2497 2502

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	y, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	p, _ := strconv.Atoi(scanner.Text())

	fmt.Println(getMinRoundsCnt(x, y, p))
}

func getMinRoundsCnt(x int, y int, p int) int {
	attackerRecord := x
	barrackRecord := y
	enemyRecord := 0

	records := make([][3]int, 1, 5000000)
	records[0] = [3]int{attackerRecord, barrackRecord, enemyRecord}

	if attackerRecord >= barrackRecord+enemyRecord {
		return 1
	}

	recordsLen := len(records)
	tmp := make([][3]int, 0, 5000000)
	roundNum := 0
	m := make(map[[3]int]struct{}, 5000000)
	for recordsLen != 0 {
		roundNum++

		for _, record := range records {
			attackerRecord = record[0]
			barrackRecord = record[1]
			enemyRecord = record[2]

			for i := 0; i <= attackerRecord; i++ {
				if i > barrackRecord {
					break
				}

				if attackerRecord-i > enemyRecord {
					continue
				}

				curBarrackRecord := barrackRecord - i
				curEnemyRecord := enemyRecord - attackerRecord + i
				curAttackerRecord := attackerRecord - curEnemyRecord
				if curAttackerRecord < 0 {
					curAttackerRecord = 0
				}

				if curBarrackRecord > 0 {
					curEnemyRecord += p
				}

				if curAttackerRecord >= curBarrackRecord+curEnemyRecord {
					return roundNum + 1
				}

				if curAttackerRecord == 0 {
					continue
				}

				diff := curEnemyRecord - curAttackerRecord
				if diff > 0 && 2*(curAttackerRecord-diff) <= diff {
					continue
				}

				if 2*curAttackerRecord <= curBarrackRecord+curEnemyRecord && curAttackerRecord <= curEnemyRecord {
					break
				}

				if attackerRecord%2 == 0 && barrackRecord%2 == 0 && enemyRecord%2 == 0 {
					attackerRecord /= 2
					barrackRecord /= 2
					enemyRecord /= 2
				}

				if attackerRecord%3 == 0 && barrackRecord%3 == 0 && enemyRecord%3 == 0 {
					attackerRecord /= 3
					barrackRecord /= 3
					enemyRecord /= 3
				}

				if attackerRecord%5 == 0 && barrackRecord%5 == 0 && enemyRecord%5 == 0 {
					attackerRecord /= 5
					barrackRecord /= 5
					enemyRecord /= 5
				}

				if attackerRecord%7 == 0 && barrackRecord%7 == 0 && enemyRecord%7 == 0 {
					attackerRecord /= 7
					barrackRecord /= 7
					enemyRecord /= 7
				}

				if _, ok := m[[3]int{
					curAttackerRecord,
					curBarrackRecord,
					curEnemyRecord,
				}]; ok {
					continue
				}
				m[[3]int{
					curAttackerRecord,
					curBarrackRecord,
					curEnemyRecord,
				}] = struct{}{}

				tmp = append(tmp, [3]int{
					curAttackerRecord,
					curBarrackRecord,
					curEnemyRecord,
				})
			}
		}

		recordsLen = len(tmp)
		records = records[0:recordsLen]
		copy(records, tmp)
		tmp = tmp[:0]
	}

	return -1
}

func getBinaryMinSearch(attackerRecord int, barrackRecord int, enemyRecord int) int {
	roundsCnt := 0

	if attackerRecord == 1 && barrackRecord == 1 && enemyRecord == 1 {
		return -1
	}

	if attackerRecord >= barrackRecord+enemyRecord {
		return 1
	}

	diff := enemyRecord - attackerRecord
	tmpBarrackRecord := barrackRecord
	tmpEnemyRecord := diff
	if diff < 0 {
		tmpBarrackRecord = barrackRecord + diff
		tmpEnemyRecord = 0
	}

	if getBinaryMinSearch(attackerRecord, tmpBarrackRecord, tmpEnemyRecord) == -1 {
		return -1
	}

	left := 0
	right := attackerRecord
	for left < right {
		middle := (left + right) / 2
		if checkResolved(middle, attackerRecord, barrackRecord, enemyRecord) {
			right = middle
		} else {
			left = middle + 1
		}
	}

	return roundsCnt
}

func checkResolved(val int, attackerRecord int, barrackRecord int, enemyRecord int) bool {
	return true
}
