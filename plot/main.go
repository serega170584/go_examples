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

	sw := [2]int{}

	scanner.Scan()
	sw[0], _ = strconv.Atoi(scanner.Text())

	scanner.Scan()
	sw[1], _ = strconv.Atoi(scanner.Text())

	ne := [2]int{}

	scanner.Scan()
	ne[0], _ = strconv.Atoi(scanner.Text())

	scanner.Scan()
	ne[1], _ = strconv.Atoi(scanner.Text())

	point := [2]int{}

	scanner.Scan()
	point[0], _ = strconv.Atoi(scanner.Text())

	scanner.Scan()
	point[1], _ = strconv.Atoi(scanner.Text())

	nw := [2]int{sw[0], ne[1]}

	se := [2]int{ne[0], sw[1]}

	xDistance, yDistance := getDistance(point, sw)
	minDistance := xDistance + yDistance
	minDistanceStr := "SW"

	xDistance, yDistance = getDistance(point, ne)
	distance := xDistance + yDistance
	if distance < minDistance {
		minDistance = distance
		minDistanceStr = "NE"
	}

	xDistance, yDistance = getDistance(point, nw)
	distance = xDistance + yDistance
	if distance < minDistance {
		minDistance = distance
		minDistanceStr = "NW"
	}

	xDistance, yDistance = getDistance(point, se)
	distance = xDistance + yDistance
	if distance < minDistance {
		minDistance = distance
		minDistanceStr = "SE"
	}

	if nw[0] <= point[0] && point[0] <= ne[0] {
		xDistance, yDistance = getDistance(point, [2]int{point[0], nw[1]})
		distance = xDistance + yDistance
		if distance < minDistance {
			minDistance = distance
			minDistanceStr = "N"
		}
	}

	if sw[1] <= point[1] && point[1] <= nw[1] {
		xDistance, yDistance = getDistance(point, [2]int{nw[0], point[1]})
		distance = xDistance + yDistance
		if distance < minDistance {
			minDistance = distance
			minDistanceStr = "W"
		}
	}

	if sw[0] <= point[0] && point[0] <= se[0] {
		xDistance, yDistance = getDistance(point, [2]int{point[0], se[1]})
		distance = xDistance + yDistance
		if distance < minDistance {
			minDistance = distance
			minDistanceStr = "S"
		}
	}

	if se[1] <= point[1] && point[1] <= ne[1] {
		xDistance, yDistance = getDistance(point, [2]int{ne[0], point[1]})
		distance = xDistance + yDistance
		if distance < minDistance {
			minDistance = distance
			minDistanceStr = "E"
		}
	}

	fmt.Println(minDistanceStr)
}

func getDistance(point [2]int, vector [2]int) (xDistance int, yDistance int) {
	xDistance = point[0] - vector[0]
	if xDistance < 0 {
		xDistance = -xDistance
	}

	yDistance = point[1] - vector[1]
	if yDistance < 0 {
		yDistance = -yDistance
	}

	return
}
