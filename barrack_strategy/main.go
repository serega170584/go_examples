package main

import (
	"bufio"
	"fmt"
	"math"
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

	diff := y - x
	if diff < 0 {
		diff = 0
	}

	existed := make(map[[3]int]int, 50000000)
	minVal := getRoundsCntMin(x, diff, 0, p, existed)
	if minVal == math.MaxInt {
		minVal = -1
	}

	fmt.Println(minVal)
}

func getRoundsCntMin(attackerRecord int, barrackRecord int, enemyRecord int, p int, existed map[[3]int]int) int {
	if v, ok := existed[[3]int{attackerRecord, barrackRecord, enemyRecord}]; ok {
		return v
	}

	baseAttackerRecord := attackerRecord
	baseBarrackRecord := barrackRecord
	baseEnemyRecord := enemyRecord

	if barrackRecord == 0 && enemyRecord == 0 {
		return 1
	}

	attackerRecord = attackerRecord - enemyRecord

	if attackerRecord <= 0 {
		return math.MaxInt
	}

	if barrackRecord > 0 {
		enemyRecord += p
	}

	if 2*attackerRecord <= barrackRecord+enemyRecord && attackerRecord <= enemyRecord {
		return math.MaxInt
	}

	minRecord := min(attackerRecord, barrackRecord)
	minRecord = min(minRecord, enemyRecord)

	last := min(attackerRecord, barrackRecord)
	startEnemyRecord := attackerRecord
	first := 0
	if attackerRecord > enemyRecord {
		first = attackerRecord - enemyRecord
	}

	if barrackRecord+enemyRecord <= attackerRecord {
		first = barrackRecord
		startEnemyRecord = barrackRecord + enemyRecord
	}

	minVal := math.MaxInt
	for i := first; i <= last; i++ {
		if attackerRecord == baseAttackerRecord && baseBarrackRecord == barrackRecord-i && baseEnemyRecord == enemyRecord-startEnemyRecord+i {
			continue
		}

		resAttackerRecord := attackerRecord
		resBarrackRecord := barrackRecord - i
		resEnemyRecord := enemyRecord - startEnemyRecord + i

		if 2*resAttackerRecord <= resBarrackRecord+resEnemyRecord && resAttackerRecord <= resEnemyRecord {
			continue
		}

		diff := resEnemyRecord - resAttackerRecord
		if diff > 0 && 2*(resAttackerRecord-diff) <= diff {
			continue
		}

		v := getRoundsCntMin(attackerRecord, barrackRecord-i, enemyRecord-startEnemyRecord+i, p, existed)
		if v == 1 {
			return v + 1
		}
		minVal = min(v, minVal)
	}

	if minVal != math.MaxInt {
		minVal += 1
	}

	existed[[3]int{baseAttackerRecord, baseBarrackRecord, baseEnemyRecord}] = minVal

	return minVal
}
