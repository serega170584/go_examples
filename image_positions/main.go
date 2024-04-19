package main

import (
	"math"
	"strconv"
	"strings"
)

func main() {

}

func getImagePositions(w int, h int, c int, s string) [][2]int {
	runes := []rune(s)

	//lf := []rune("\n")[0]
	space := []rune(" ")[0]
	leftBracket := []rune("(")[0]
	rightBracket := []rune(")")[0]

	//isEmptyString := false
	isImage := false
	image := make([]rune, 0)

	fragments := make([][3]int, 0)
	fragments = append(fragments, [3]int{0, w, math.MaxInt})
	insertFragment := [3]int{0, 0, h}
	columnPosition := 0
	rowPosition := 0
	currentFragmentInd := 0
	prevRowHeight := 0
	floatFragment := [5]int{0, 0, 0, 0, 0}
	cornerRowPosition := 0

	imagePositions := make([][2]int, 0)
	isFragmentImage := false

	for _, v := range runes {
		if !isImage {
			insertFragment[1] += c
		}

		isImageHandled := false
		if isImage {
			image = append(image, v)
			if v == rightBracket {
				isImage = false
				paramsStrings := strings.Split(string(image), " ")
				paramsStringsLength := len(paramsStrings)
				params := make(map[string]string, paramsStringsLength)
				for _, paramsString := range paramsStrings {
					pair := strings.Split(paramsString, "=")
					params[pair[0]] = pair[1]
				}
				width, _ := strconv.Atoi(params["width"])
				height, _ := strconv.Atoi(params["height"])
				layout, _ := params["layout"]
				insertFragment = [3]int{0, 0, 0}
				floatFragment = [5]int{0, 0, 0, 0, 0}
				if layout == "embedded" || layout == "surrounded" {
					insertFragment = [3]int{0, width, height}
					isFragmentImage = true
				} else {
					dx, _ := strconv.Atoi(params["dx"])
					dy, _ := strconv.Atoi(params["dy"])
					xPosition := columnPosition + dx
					yPosition := rowPosition + dy
					if xPosition < 0 {
						xPosition = 0
					}
					if yPosition < 0 {
						yPosition = 0
					}
					if xPosition+width > w {
						xPosition = w - width
					}
					floatFragment = [5]int{0, width, height, xPosition, yPosition}
				}
				isImageHandled = true
				if layout == "embedded" {
					prevRowHeight = max(prevRowHeight, height)
				}
			} else {
				continue
			}
		}

		if v == space || isImageHandled {
			if floatFragment != [5]int{0, 0, 0, 0, 0} {
				imagePosition := [2]int{floatFragment[3], floatFragment[4]}
				imagePositions = append(imagePositions, imagePosition)
				floatFragment = [5]int{0, 0, 0, 0, 0}
			}

			if insertFragment == [3]int{0, 0, h} {
				continue
			}

			isFilled := false
			for !isFilled {
				fragmentsForInsert := make([][3]int, 0)
				for i := currentFragmentInd; i < len(fragments); i++ {
					if fragments[i][2] > 0 && insertFragment[1] <= fragments[i][1]-fragments[i][0]+columnPosition {
						if isFragmentImage {
							imagePositions = append(imagePositions, [2]int{columnPosition, rowPosition})
							isFragmentImage = false
							cornerRowPosition = max(cornerRowPosition, rowPosition+insertFragment[2])
							fragmentForInsert := [3]int{columnPosition - fragments[i][0], fragments[i][1], fragments[i][2]}
							fragmentsForInsert = append(fragmentsForInsert, fragmentForInsert)
							fragmentForInsert = [3]int{columnPosition, insertFragment[1], -insertFragment[2]}
							fragmentsForInsert = append(fragmentsForInsert, fragmentForInsert)
							fragmentForInsert = [3]int{columnPosition, fragments[i][0] - insertFragment[1], fragments[i][2]}
							fragmentsForInsert = append(fragmentsForInsert, fragmentForInsert)
							currentFragmentInd = i + 2
						} else {
							currentFragmentInd = i
						}
						columnPosition += insertFragment[1]
						isFilled = true
						break
					}
					columnPosition += fragments[currentFragmentInd][1]
				}

				if !isFilled {
					rowPosition += prevRowHeight
					prevRowHeight = h
					oldFragments := make([][3]int, 0)
					fragment := [3]int{0, 0, math.MaxInt}
					columnPosition = 0
					for i := 0; i < len(fragments); i++ {
						if fragments[i][2] < 0 {
							fragments[i][2] += prevRowHeight
						}

						if fragments[i][2] >= 0 {
							if fragment[2] < 0 {
								oldFragments = append(oldFragments, fragment)
								fragment[0] = fragments[i][0]
								fragment[1] = fragments[i][1]
								fragment[2] = math.MaxInt
							} else {
								fragment[1] += fragments[i][1]
							}
						}

						if fragments[i][2] < 0 {
							if fragment != [3]int{0, 0, math.MaxInt} {
								oldFragments = append(oldFragments, fragment)
							}
							fragment = [3]int{fragments[i][0], fragments[i][1], fragments[i][2]}
						}
					}

					oldFragments = append(oldFragments, fragment)

					fragments = make([][3]int, len(oldFragments))
					copy(fragments, oldFragments)
				}
			}
		}

		if v == leftBracket {
			isImage = true
			image = make([]rune, 0)
			image = append(image, v)
			continue
		}

		//
		//if v != space && v != lf {
		//	word = append(word, v)
		//	wordSymPosition++
		//	isEmptyString = true
		//}
		//
		//if v == space && prevSym == lf {
		//	isEmptyString = true
		//}
		//
		//if v == space {
		//	pixelXPosition += c * wordSymPosition
		//	wordSymPosition = 0
		//}
		//
		//if v == lf && isEmptyString {
		//	pixelYPosition += h
		//}
		//
		//if v == lf {
		//	pixelXPosition = 0
		//}
		//
		//prevSym = v
	}
}
