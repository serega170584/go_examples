package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// 2500 5000 2499
// 2500 4000 0
// 2500 1500 2499
// 2500 0 1499
// 1001 0 1499
// 1001 0 498
// 503 0 498

// 2500 4100 0
// 2500 1600 2499
// 2500 0 1599
// 901 0 1599
// 901 0 698
// 203 0 698
// 203 0

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

	existed := make(map[[3]int32]int32, 500)
	minVal := getRoundsCntMin(x, diff, 0, p, existed)
	if minVal == math.MaxInt {
		minVal = -1
	}

	fmt.Println(minVal)
}

func getRoundsCntMin(attackerRecord int, barrackRecord int, enemyRecord int, p int, existed map[[3]int32]int32) int {
	if v, ok := existed[[3]int32{int32(attackerRecord), int32(barrackRecord), int32(enemyRecord)}]; ok {
		return int(v)
	}

	if attackerRecord+enemyRecord < barrackRecord && enemyRecord == 0 && attackerRecord-p > 0 {
		cnt := barrackRecord / (attackerRecord + enemyRecord)
		minVal := getRoundsCntMin(attackerRecord, barrackRecord-cnt*(attackerRecord-p), 0, p, existed) + cnt
		existed[[3]int32{int32(attackerRecord), int32(attackerRecord), int32(attackerRecord)}] = int32(minVal)
		return minVal
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

	if 2*attackerRecord <= barrackRecord+p && attackerRecord <= p && barrackRecord != 0 {
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
	if !(attackerRecord == baseAttackerRecord && baseBarrackRecord == barrackRecord-first && baseEnemyRecord == enemyRecord-startEnemyRecord+first) {
		minVal = getRoundsCntMin(attackerRecord, barrackRecord-first, enemyRecord-startEnemyRecord+first, p, existed)
		if first == last {
			if minVal == math.MaxInt {
				return minVal
			}
			return minVal + 1
		}
		if minVal == 1 {
			return 2
		}
	}

	if !(attackerRecord == baseAttackerRecord && baseBarrackRecord == barrackRecord-last && baseEnemyRecord == enemyRecord-startEnemyRecord+last) {
		v := getRoundsCntMin(attackerRecord, barrackRecord-last, enemyRecord-startEnemyRecord+last, p, existed)
		minVal = min(v, minVal)
	}

	if minVal != math.MaxInt {
		minVal += 1
	}

	existed[[3]int32{int32(baseAttackerRecord), int32(baseBarrackRecord), int32(baseEnemyRecord)}] = int32(minVal)

	return minVal
}
