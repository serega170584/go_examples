package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	l := len(ranked)
	uniq := make([]int32, 0, l)
	prev := int32(-1)
	for _, v := range ranked {
		if prev != v {
			uniq = append(uniq, v)
		}
		prev = v
	}

	i := len(uniq) - 1
	playerRanks := make([]int32, 0, len(player))
	pi := -1
	for j, p := range player {
		for i >= 0 && uniq[i] < p {
			i--
		}

		pi = j

		if i == -1 {
			break
		}

		rank := int32(i)
		if uniq[i] > p {
			rank++
		}

		playerRanks = append(playerRanks, rank+1)
	}

	pl := len(player)
	for k := 0; k < pl-pi; k++ {
		playerRanks = append(playerRanks, 1)
	}

	return playerRanks
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout := os.Stdout
	//checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	rankedCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	rankedTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ranked []int32

	for i := 0; i < int(rankedCount); i++ {
		rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
		checkError(err)
		rankedItem := int32(rankedItemTemp)
		ranked = append(ranked, rankedItem)
	}

	playerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	playerTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var player []int32

	for i := 0; i < int(playerCount); i++ {
		playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
		checkError(err)
		playerItem := int32(playerItemTemp)
		player = append(player, playerItem)
	}

	result := climbingLeaderboard(ranked, player)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
