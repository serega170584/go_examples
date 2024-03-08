package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 2 2 1  2 1 0  2 1 1
// 2500 5000 2499
// 2500 2500 2499
// 2500 2499 2499
// 2500 2498 2499

// 1250 4750 1249
// 250 3750 249  3351
// 50 3550 49  3473
// 2500 5000 2498
// 1250 2500 1249
// 625 1250 624
// 40 44 60

// 626 1250 624 124
// 313 625 312   124

// 10 5000 2  625
// 9 4999 1

// 500 3000 499
//

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

			// attackerRecord - i

			// 2 2 2

			// attackerRecord - i <= 2 * enemyRecord
			//

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

				if 2*curAttackerRecord <= curBarrackRecord+curEnemyRecord && curAttackerRecord == curEnemyRecord {
					continue
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
