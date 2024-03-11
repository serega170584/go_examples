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
	holidayCnt, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	year, _ := strconv.Atoi(scanner.Text())

	holidayDays := make([]string, holidayCnt)
	for i := 0; i < holidayCnt; i++ {
		str := make([]string, 2)
		scanner.Scan()
		str[0] = scanner.Text()
		scanner.Scan()
		str[1] = scanner.Text()
		holidayDays[i] = strings.Join(str, " ")
	}

	scanner.Scan()
	firstYearDay := scanner.Text()

	fmt.Println(specialHolidayDays(holidayCnt, year, holidayDays, firstYearDay))
}

func specialHolidayDays(holidayCnt int, year int, holidayDays []string, firstYearDay string) (string, string) {
	weekDays := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	weekDaysMap := make(map[string]int, 7)
	weekDaysMap["Monday"] = 0
	weekDaysMap["Tuesday"] = 1
	weekDaysMap["Wednesday"] = 2
	weekDaysMap["Thursday"] = 3
	weekDaysMap["Friday"] = 4
	weekDaysMap["Saturday"] = 5
	weekDaysMap["Sunday"] = 6

	isDifficultYear := false
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		isDifficultYear = true
	}

	monthDayCounts := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if isDifficultYear {
		monthDayCounts[1] = 29
	}

	startMonthWeekDays := make([]int, 12)
	startMonthWeekDays[0] = weekDaysMap[firstYearDay] // 3
	for i := 1; i < 12; i++ {
		weekDayOrder := startMonthWeekDays[i-1] + monthDayCounts[i-1]%7
		if weekDayOrder > 6 {
			weekDayOrder = weekDayOrder - 7
		}
		startMonthWeekDays[i] = weekDayOrder
	}

	monthsMap := make(map[string]int)
	monthsMap["January"] = 0
	monthsMap["February"] = 1
	monthsMap["March"] = 2
	monthsMap["April"] = 3
	monthsMap["May"] = 4
	monthsMap["June"] = 5
	monthsMap["July"] = 6
	monthsMap["August"] = 7
	monthsMap["September"] = 8
	monthsMap["October"] = 9
	monthsMap["November"] = 10
	monthsMap["December"] = 11

	holidayCounts := make([]int, 7)
	for _, val := range holidayDays {
		parts := strings.Split(val, " ")
		startDay := startMonthWeekDays[monthsMap[parts[1]]]
		day, _ := strconv.Atoi(parts[0])
		day = startDay + (day-1)%7
		if day > 6 {
			day = day - 7
		}
		holidayCounts[day]++
	}

	yearHolidayCounts := make([]int, 7)
	for i := 0; i < 7; i++ {
		yearHolidayCounts[i] = 52
	}

	yearHolidayCounts[weekDaysMap[firstYearDay]] = 53
	if isDifficultYear {
		curDay := weekDaysMap[firstYearDay] + 1
		if curDay > 6 {
			curDay = curDay - 7
		}
		yearHolidayCounts[curDay] = 53
	}

	worst := ""
	best := ""
	minVal := 52 + holidayCnt
	maxVal := 52
	for i := range yearHolidayCounts {
		yearHolidayCounts[i] += holidayCnt - holidayCounts[i]
		minVal = min(minVal, yearHolidayCounts[i])
		if minVal == yearHolidayCounts[i] {
			worst = weekDays[i]
		}
		maxVal = max(maxVal, yearHolidayCounts[i])
		if maxVal == yearHolidayCounts[i] {
			best = weekDays[i]
		}
	}

	return best, worst
}
