package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Graph struct {
	vertexesList [][]int
	marked       []bool
}

func NewGraph(cnt int, edges [][2]int) *Graph {
	vertexesList := make([][]int, cnt)
	for i := range vertexesList {
		vertexesList[i] = make([]int, cnt)
	}

	for _, val := range edges {
		firstVertex := val[0]
		vertexesList[firstVertex] = append(vertexesList[firstVertex], val[1])
		secondVertex := val[1]
		vertexesList[secondVertex] = append(vertexesList[secondVertex], val[0])
	}

	marked := make([]bool, cnt)

	return &Graph{marked: marked, vertexesList: vertexesList}
}

func (g *Graph) dfs(v int) {
	g.marked[v] = true
	for _, val := range g.vertexesList[v] {
		if !g.marked[val] {
			g.dfs(val)
		}
	}
	fmt.Println(v)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter vertexes count")
	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter edges count")
	scanner.Scan()
	edgesCnt, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter edges")
	edges := make([][2]int, edgesCnt)
	for i := 0; i < edgesCnt; i++ {
		scanner.Scan()
		edges[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		edges[i][1], _ = strconv.Atoi(scanner.Text())
	}

	g := NewGraph(cnt, edges)
	g.dfs(0)
}
