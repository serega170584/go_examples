package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Graph struct {
	vertexesList [][]int
	vertexesCnt  int
	marked       []bool
}

func NewGraph(vertexesCnt int, cnt int, edges [][2]int) *Graph {
	vertexesList := make([][]int, vertexesCnt)
	marked := make([]bool, vertexesCnt)
	for i := 0; i < cnt; i++ {
		firstVertex := edges[i][0]
		secondVertex := edges[i][1]
		vertexesList[firstVertex] = append(vertexesList[firstVertex], secondVertex)
		vertexesList[secondVertex] = append(vertexesList[firstVertex], firstVertex)
	}
	return &Graph{
		vertexesList: vertexesList,
		vertexesCnt:  vertexesCnt,
		marked:       marked,
	}
}

func (g *Graph) dfs(i int) {
	g.marked[i] = true
	neigbours := g.vertexesList[i]
	for _, v := range neigbours {
		if !g.marked[v] {
			g.dfs(v)
		}
	}
	fmt.Println(i)
}

func (g *Graph) calculateComponentsCount() int {
	cnt := 0
	for i := range g.vertexesList {
		if !g.marked[i] {
			g.dfs(i)
			cnt++
		}
	}
	return cnt
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter vertexes count")
	scanner.Scan()
	vertexesCnt, _ := strconv.Atoi(scanner.Text())

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

	g := NewGraph(vertexesCnt, edgesCnt, edges)
	fmt.Println("Got components count", g.calculateComponentsCount())
}
