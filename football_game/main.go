package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("football_game/123")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	text := make([]string, 0, 500)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		if scanner.Text() == "" {
			break
		} else {
			text = append(text, scanner.Text())
		}
	}
	//re := regexp.MustCompile(`^([ a-zA-Z]*) ([0-9]*[0-9]*)'$`)
	//fmt.Println(re.FindStringSubmatch(`Del Piero 67'`))
	//scanner := makeScanner()
	//for scanner.Scan() {
	//	if scanner.Text() == "" {
	//		break
	//	} else {
	//		text = append(text, scanner.Text())
	//	}
	//}
	stat := getStat(text)
	for _, v := range stat {
		fmt.Println(v)
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func getStat(text []string) []float64 {
	playRe := regexp.MustCompile(`^"([a-zA-Z ]*)" - "([a-zA-Z ]*)" ([0-9]*[0-9])*:([0-9]*[0-9]*)$`)
	goalsRe := regexp.MustCompile(`^([ a-zA-Z]*) ([0-9]*[0-9]*)'$`)

	totalGoalsReportRe := regexp.MustCompile(`^Total goals for "([a-zA-Z ]*)"$`)
	meanGoalsReportRe := regexp.MustCompile(`^Mean goals per game for "([a-zA-Z ]*)"$`)
	playerTotalGoalsReportRe := regexp.MustCompile(`^Total goals by ([a-zA-Z ]*)$`)
	playerMeanGoalsReportRe := regexp.MustCompile(`^Mean goals per game by ([a-zA-Z ]*)$`)
	playerMinuteGoalsReportRe := regexp.MustCompile(`^Goals on minute ([0-9]*[0-9]*) by ([a-zA-Z ]*)$`)
	playerFirstMinutesGoalsReportRe := regexp.MustCompile(`^Goals on first ([0-9]*[0-9]*) minutes by ([a-zA-Z ]*)$`)
	playerLastMinutesGoalsReportRe := regexp.MustCompile(`^Goals on last ([0-9]*[0-9]*) minutes by ([a-zA-Z ]*)$`)
	teamOpensReportRe := regexp.MustCompile(`^Score opens by "([a-zA-Z ]*)"$`)
	playerOpensReportRe := regexp.MustCompile(`^Score opens by ([a-zA-Z ]*)$`)

	teamGoalsCnt := make(map[string]int, 100)
	teamPlaysCnt := make(map[string]int, 100)

	playerGoalsCnt := make(map[string]int, 200)

	firstTeam := ""
	secondTeam := ""
	goalNumber := 0

	playerTeam := make(map[string]string, 200)

	playerGoalMinutesCnt := make(map[string]map[int]int, 200)

	openMinute := math.MaxInt

	playerOpens := make(map[string]int, 200)
	teamOpens := make(map[string]int, 200)

	openPlayer := ""

	firstTeamGoalsCnt := 0
	secondTeamGoalsCnt := 0

	reportItems := make([]float64, 0, 500)

	for _, v := range text {
		playInfo := playRe.FindStringSubmatch(v)
		goalsInfo := goalsRe.FindStringSubmatch(v)
		totalGoalsReportInfo := totalGoalsReportRe.FindStringSubmatch(v)
		meanGoalsReportInfo := meanGoalsReportRe.FindStringSubmatch(v)
		playerTotalGoalsReportInfo := playerTotalGoalsReportRe.FindStringSubmatch(v)
		playerMeanGoalsReportInfo := playerMeanGoalsReportRe.FindStringSubmatch(v)
		playerMinuteGoalsReportInfo := playerMinuteGoalsReportRe.FindStringSubmatch(v)
		playerFirstMinutesGoalsReportInfo := playerFirstMinutesGoalsReportRe.FindStringSubmatch(v)
		playerLastMinutesGoalsReportInfo := playerLastMinutesGoalsReportRe.FindStringSubmatch(v)
		teamOpensReportInfo := teamOpensReportRe.FindStringSubmatch(v)
		playerOpensReportInfo := playerOpensReportRe.FindStringSubmatch(v)
		if len(playInfo) != 0 {
			firstTeam = playInfo[1]
			secondTeam = playInfo[2]

			firstTeamGoalsCnt, _ = strconv.Atoi(playInfo[3])
			secondTeamGoalsCnt, _ = strconv.Atoi(playInfo[4])

			teamGoalsCnt[firstTeam] += firstTeamGoalsCnt
			teamGoalsCnt[secondTeam] += secondTeamGoalsCnt

			teamPlaysCnt[firstTeam]++
			teamPlaysCnt[secondTeam]++

			goalNumber = 0

			openMinute = math.MaxInt

			openPlayer = ""
		} else if len(goalsInfo) != 0 {
			player := goalsInfo[1]
			minute, _ := strconv.Atoi(goalsInfo[2])

			if goalNumber < firstTeamGoalsCnt {
				playerTeam[player] = firstTeam
			} else {
				playerTeam[player] = secondTeam
			}

			if minute < openMinute {
				openMinute = minute
				openPlayer = player
			}

			if playerGoalMinutesCnt[player] == nil {
				playerGoalMinutesCnt[player] = make(map[int]int, 90)
			}
			playerGoalMinutesCnt[player][minute]++
			playerGoalsCnt[player]++
			goalNumber++

			if goalNumber == firstTeamGoalsCnt+secondTeamGoalsCnt {
				if openPlayer != "" {
					playerOpens[openPlayer]++
					teamOpens[playerTeam[openPlayer]]++
				}
			}
		} else if len(totalGoalsReportInfo) != 0 {
			team := totalGoalsReportInfo[1]
			cnt := 0
			if _, ok := teamGoalsCnt[team]; ok {
				cnt = teamGoalsCnt[team]
			}
			reportItems = append(reportItems, float64(cnt))
		} else if len(meanGoalsReportInfo) != 0 {
			team := meanGoalsReportInfo[1]
			goalsCnt := 0
			cnt := 0
			if _, ok := teamGoalsCnt[team]; ok {
				goalsCnt = teamGoalsCnt[team]
			}
			if _, ok := teamPlaysCnt[team]; ok {
				cnt = teamPlaysCnt[team]
			}
			res := float64(0)
			if cnt != 0 {
				res = float64(goalsCnt) / float64(cnt)
			}
			reportItems = append(reportItems, res)
		} else if len(playerTotalGoalsReportInfo) != 0 {
			player := playerTotalGoalsReportInfo[1]
			cnt := 0
			if _, ok := playerGoalsCnt[player]; ok {
				cnt = playerGoalsCnt[player]
			}
			reportItems = append(reportItems, float64(cnt))
		} else if len(playerMeanGoalsReportInfo) != 0 {
			player := playerMeanGoalsReportInfo[1]
			goalsCnt := 0
			cnt := 0
			if _, ok := playerGoalsCnt[player]; ok {
				goalsCnt = playerGoalsCnt[player]
			}
			if _, ok := teamPlaysCnt[playerTeam[player]]; ok {
				cnt = teamPlaysCnt[playerTeam[player]]
			}
			res := float64(0)
			if cnt != 0 {
				res = float64(goalsCnt) / float64(cnt)
			}
			reportItems = append(reportItems, res)
		} else if len(playerMinuteGoalsReportInfo) != 0 {
			player := playerMinuteGoalsReportInfo[2]
			minute, _ := strconv.Atoi(playerMinuteGoalsReportInfo[1])
			cnt := 0
			if _, ok := playerGoalMinutesCnt[player][minute]; ok {
				cnt = playerGoalMinutesCnt[player][minute]
			}
			reportItems = append(reportItems, float64(cnt))
		} else if len(playerFirstMinutesGoalsReportInfo) != 0 {
			player := playerFirstMinutesGoalsReportInfo[2]
			minute, _ := strconv.Atoi(playerFirstMinutesGoalsReportInfo[1])
			cnt := 0
			for m, c := range playerGoalMinutesCnt[player] {
				if m <= minute {
					cnt += c
				}
			}
			reportItems = append(reportItems, float64(cnt))
		} else if len(playerLastMinutesGoalsReportInfo) != 0 {
			player := playerLastMinutesGoalsReportInfo[2]
			minute, _ := strconv.Atoi(playerLastMinutesGoalsReportInfo[1])
			minute = 91 - minute
			cnt := 0
			for m, c := range playerGoalMinutesCnt[player] {
				if 90 >= m && m >= minute {
					cnt += c
				}
			}
			reportItems = append(reportItems, float64(cnt))
		} else if len(teamOpensReportInfo) != 0 {
			team := teamOpensReportInfo[1]
			cnt := 0
			if _, ok := teamOpens[team]; ok {
				cnt = teamOpens[team]
			}
			reportItems = append(reportItems, float64(cnt))
		} else if len(playerOpensReportInfo) != 0 {
			player := playerOpensReportInfo[1]
			cnt := 0
			if _, ok := playerOpens[player]; ok {
				cnt = playerOpens[player]
			}
			reportItems = append(reportItems, float64(cnt))
		}
	}

	return reportItems
}
