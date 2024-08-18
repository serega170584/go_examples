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

	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	counts := make([]int, cnt)
	list := make([]string, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		counts[i], _ = strconv.Atoi(scanner.Text())

		scanner.Scan()
		list[i] = scanner.Text()
	}

	resCnt, playlists := getPlaylists(list, counts, cnt)

	fmt.Println(resCnt)
	fmt.Println(playlists)
}

func getPlaylists(list []string, counts []int, cnt int) (int, []string) {
	names := make(map[string]int)
	for i, v := range list {
		nv := strings.Split(v, " ")
		for j := 0; j < counts[i]; j++ {
			names[nv[j]]++
		}
	}

	playLists := make([]string, 0)
	resCnt := 0
	for name, v := range names {
		if v == cnt {
			playLists = append(playLists, name)
			resCnt++
		}
	}

	return resCnt, playLists
}
