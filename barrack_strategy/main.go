package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

			minVal := min(barrackRecord, enemyRecord)
			minVal = min(minVal, attackerRecord)
			first := 0
			last := attackerRecord
			if minVal == enemyRecord {
				first = min(barrackRecord, attackerRecord) - enemyRecord
				if attackerRecord > barrackRecord {
					last = barrackRecord
				}
			}
			if minVal == barrackRecord {
				last = barrackRecord
				if attackerRecord-enemyRecord > 0 {
					first += attackerRecord - enemyRecord
				}
			}

			//if last >= barrackRecord {
			//	curBarrackRecord := 0
			//	curEnemyRecord := enemyRecord - attackerRecord + barrackRecord
			//	curAttackerRecord := attackerRecord - curEnemyRecord
			//
			//	if curAttackerRecord < 0 {
			//		curAttackerRecord = 0
			//	}
			//
			//	tmp = append(tmp, [3]int{curAttackerRecord, curBarrackRecord, curEnemyRecord})
			//}

			for i := first; i <= last; i++ {
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
