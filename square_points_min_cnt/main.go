package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	points := make([][2]int, n)
	for i := 0; i < n; i++ {
		point := [2]int{}
		scanner.Scan()
		point[0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		point[1], _ = strconv.Atoi(scanner.Text())
		points[i] = point
	}

	cnt, neededPoints := getSquarePointsMinCnt(n, points)

	fmt.Println(cnt)
	if cnt != 0 {
		for _, v := range neededPoints {
			str := []string{strconv.Itoa(v[0]), strconv.Itoa(v[1])}
			fmt.Println(strings.Join(str, " "))
		}
	}
}

func getSquarePointsMinCnt(n int, points [][2]int) (int, [][2]int) {
	squarePointCnts := make(map[[4]int]int, n*(n+1))

	markedPoints := make(map[int]struct{}, n+1)
	pointsIndexesMap := make(map[[2]int]int, n*(n+1))
	pointsMap := make(map[int][2]int, n*(n+1))
	point := 0
	pointsIndexesMap[[2]int{100000001, 100000001}] = 0
	pointsMap[0] = [2]int{100000001, 100000001}
	markedPoints[0] = struct{}{}
	lastPointIndex := 1
	for _, v := range points {
		pointsIndexesMap[v] = lastPointIndex
		pointsMap[lastPointIndex] = v
		curVInd := lastPointIndex
		lastPointIndex++
		for i := range markedPoints {
			p := pointsMap[i]
			firstPointX := v[0]
			firstPointY := v[1]
			secondPointX := p[0]
			secondPointY := p[1]

			movX := -firstPointX
			firstPointX = 0
			secondPointX += movX

			movY := -secondPointY
			firstPointY += movY
			secondPointY = 0

			startOrtFirstX := 0
			startOrtFirstY := secondPointX

			startOrtSecondX := -firstPointY
			startOrtSecondY := 0

			xDiff := firstPointX - startOrtFirstX
			yDiff := firstPointY - startOrtFirstY

			firstSideThirdPointX := startOrtSecondX + xDiff
			firstSideThirdPointY := startOrtSecondY + yDiff

			xDiff = secondPointX - firstPointX
			yDiff = secondPointY - firstPointY

			firstSideFourthPointX := firstSideThirdPointX + xDiff
			firstSideFourthPointY := firstSideThirdPointY + yDiff

			xDiff = 2 * (firstPointX - firstSideThirdPointX)
			yDiff = 2 * (firstPointY - firstSideThirdPointY)

			secondSideThirdPointX := firstSideThirdPointX + xDiff
			secondSideThirdPointY := firstSideThirdPointY + yDiff

			xDiff = 2 * (secondPointX - firstSideFourthPointX)
			yDiff = 2 * (secondPointY - firstSideFourthPointY)

			secondSideFourthPointX := firstSideFourthPointX + xDiff
			secondSideFourthPointY := firstSideFourthPointY + yDiff

			firstPointX -= movX
			firstPointY -= movY

			secondPointX -= movX
			secondPointY -= movY

			firstSideThirdPointX -= movX
			firstSideThirdPointY -= movY

			firstSideFourthPointX -= movX
			firstSideFourthPointY -= movY

			secondSideThirdPointX -= movX
			secondSideThirdPointY -= movY

			secondSideFourthPointX -= movX
			secondSideFourthPointY -= movY

			curPointsList := [4]int{}
			curPointInd := -1
			if ind, ok := pointsIndexesMap[[2]int{firstPointX, firstPointY}]; !ok {
				pointsIndexesMap[[2]int{firstPointX, firstPointY}] = lastPointIndex
				pointsMap[lastPointIndex] = [2]int{firstPointX, firstPointY}
				curPointInd = lastPointIndex
				curVInd = lastPointIndex
				lastPointIndex++
			} else {
				curPointInd = ind
			}
			curPointsList[0] = curPointInd

			if ind, ok := pointsIndexesMap[[2]int{secondPointX, secondPointY}]; !ok {
				pointsIndexesMap[[2]int{secondPointX, secondPointY}] = lastPointIndex
				pointsMap[lastPointIndex] = [2]int{secondPointX, secondPointY}
				curPointInd = lastPointIndex
				lastPointIndex++
			} else {
				curPointInd = ind
			}
			curPointsList[1] = curPointInd

			if ind, ok := pointsIndexesMap[[2]int{firstSideThirdPointX, firstSideThirdPointY}]; !ok {
				pointsIndexesMap[[2]int{firstSideThirdPointX, firstSideThirdPointY}] = lastPointIndex
				pointsMap[lastPointIndex] = [2]int{firstSideThirdPointX, firstSideThirdPointY}
				curPointInd = lastPointIndex
				lastPointIndex++
			} else {
				curPointInd = ind
			}
			curPointsList[2] = curPointInd

			if ind, ok := pointsIndexesMap[[2]int{firstSideFourthPointX, firstSideFourthPointY}]; !ok {
				pointsIndexesMap[[2]int{firstSideFourthPointX, firstSideFourthPointY}] = lastPointIndex
				pointsMap[lastPointIndex] = [2]int{firstSideFourthPointX, firstSideFourthPointY}
				curPointInd = lastPointIndex
				lastPointIndex++
			} else {
				curPointInd = ind
			}
			curPointsList[3] = curPointInd

			curPointsSlice := curPointsList[:]
			slices.Sort(curPointsSlice)

			if _, ok := squarePointCnts[curPointsList]; !ok {
				squarePointCnts[curPointsList] = 4
			}

			secondCurPointsList := [4]int{}
			secondCurPointInd := -1
			if ind, ok := pointsIndexesMap[[2]int{firstPointX, firstPointY}]; !ok {
				pointsIndexesMap[[2]int{firstPointX, firstPointY}] = lastPointIndex
				pointsMap[lastPointIndex] = [2]int{firstPointX, firstPointY}
				secondCurPointInd = lastPointIndex
				lastPointIndex++
			} else {
				secondCurPointInd = ind
			}
			secondCurPointsList[0] = secondCurPointInd

			if ind, ok := pointsIndexesMap[[2]int{secondPointX, secondPointY}]; !ok {
				pointsIndexesMap[[2]int{secondPointX, secondPointY}] = lastPointIndex
				pointsMap[lastPointIndex] = [2]int{secondPointX, secondPointY}
				secondCurPointInd = lastPointIndex
				lastPointIndex++
			} else {
				secondCurPointInd = ind
			}
			secondCurPointsList[1] = secondCurPointInd

			if ind, ok := pointsIndexesMap[[2]int{secondSideThirdPointX, secondSideThirdPointY}]; !ok {
				pointsIndexesMap[[2]int{secondSideThirdPointX, secondSideThirdPointY}] = lastPointIndex
				pointsMap[lastPointIndex] = [2]int{secondSideThirdPointX, secondSideThirdPointY}
				secondCurPointInd = lastPointIndex
				lastPointIndex++
			} else {
				secondCurPointInd = ind
			}
			secondCurPointsList[2] = secondCurPointInd

			if ind, ok := pointsIndexesMap[[2]int{secondSideFourthPointX, secondSideFourthPointY}]; !ok {
				pointsIndexesMap[[2]int{secondSideFourthPointX, secondSideFourthPointY}] = lastPointIndex
				pointsMap[lastPointIndex] = [2]int{secondSideFourthPointX, secondSideFourthPointY}
				secondCurPointInd = lastPointIndex
				lastPointIndex++
			} else {
				secondCurPointInd = ind
			}
			secondCurPointsList[3] = secondCurPointInd

			secondCurPointsSlice := secondCurPointsList[:]
			slices.Sort(secondCurPointsSlice)

			if _, ok := squarePointCnts[secondCurPointsList]; !ok {
				squarePointCnts[secondCurPointsList] = 4
			}
		}
		markedPoints[curVInd] = struct{}{}
	}

	// 0 2 -> 0
	// 2 0 -> 1
	// 2 2 -> 6
	// 0 4 -> 7
	minCnt := 4
	searchedKey := [4]int{}
	for pis, c := range squarePointCnts {
		count := c
		for _, v := range pis {
			if v == point {
				continue
			}
			if _, ok := markedPoints[v]; ok {
				count--
			}
		}
		squarePointCnts[pis] = count
		if count < minCnt {
			minCnt = count
			searchedKey = pis
		}
	}

	neededPoints := make([][2]int, 0, 4)
	for _, p := range searchedKey {
		if p == point {
			neededPoints = append(neededPoints, pointsMap[p])
			continue
		}
		if _, ok := markedPoints[p]; !ok {
			neededPoints = append(neededPoints, pointsMap[p])
		}
	}

	return minCnt, neededPoints
}
