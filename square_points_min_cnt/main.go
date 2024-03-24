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
	n, _ := strconv.Atoi(scanner.Text())

	points := make(map[[2]int]struct{}, n)
	for i := 0; i < n; i++ {
		point := [2]int{}
		scanner.Scan()
		point[0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		point[1], _ = strconv.Atoi(scanner.Text())
		points[point] = struct{}{}
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

func getSquarePointsMinCnt(n int, points map[[2]int]struct{}) (int, [][2]int) {
	minCnt := 4
	totalNeededPoints := make([][2]int, 0, 4)
	for i, _ := range points {
		for p, _ := range points {
			firstPointX := i[0]
			firstPointY := i[1]
			secondPointX := p[0]
			secondPointY := p[1]

			if p == i {
				secondPointX = firstPointX + 1
				secondPointY = firstPointY
			}

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

			cnt := 4
			neededPoints := make([][2]int, 0, 4)

			if _, ok := points[[2]int{firstPointX, firstPointY}]; ok {
				cnt--
			} else {
				neededPoints = append(neededPoints, [2]int{firstPointX, firstPointY})
			}

			if _, ok := points[[2]int{secondPointX, secondPointY}]; ok {
				cnt--
			} else {
				neededPoints = append(neededPoints, [2]int{secondPointX, secondPointY})
			}

			if _, ok := points[[2]int{firstSideThirdPointX, firstSideThirdPointY}]; ok {
				cnt--
			} else {
				neededPoints = append(neededPoints, [2]int{firstSideThirdPointX, firstSideThirdPointY})
			}

			if _, ok := points[[2]int{firstSideFourthPointX, firstSideFourthPointY}]; ok {
				cnt--
			} else {
				neededPoints = append(neededPoints, [2]int{firstSideFourthPointX, firstSideFourthPointY})
			}

			if cnt < minCnt {
				minCnt = cnt
				totalNeededPoints = totalNeededPoints[:len(neededPoints)]
				totalNeededPoints = neededPoints
			}

			cnt = 4
			secondNeededPoints := make([][2]int, 0, 4)

			if _, ok := points[[2]int{firstPointX, firstPointY}]; ok {
				cnt--
			} else {
				secondNeededPoints = append(secondNeededPoints, [2]int{firstPointX, firstPointY})
			}

			if _, ok := points[[2]int{secondPointX, secondPointY}]; ok {
				cnt--
			} else {
				secondNeededPoints = append(secondNeededPoints, [2]int{secondPointX, secondPointY})
			}

			if _, ok := points[[2]int{secondSideThirdPointX, secondSideThirdPointY}]; ok {
				cnt--
			} else {
				secondNeededPoints = append(secondNeededPoints, [2]int{secondSideThirdPointX, secondSideThirdPointY})
			}

			if _, ok := points[[2]int{secondSideFourthPointX, secondSideFourthPointY}]; ok {
				cnt--
			} else {
				secondNeededPoints = append(secondNeededPoints, [2]int{secondSideFourthPointX, secondSideFourthPointY})
			}

			if cnt < minCnt {
				minCnt = cnt
				totalNeededPoints = totalNeededPoints[:len(secondNeededPoints)]
				totalNeededPoints = secondNeededPoints
			}
		}

	}

	return minCnt, totalNeededPoints
}
