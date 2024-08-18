package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Graph struct {
	v   int
	e   int
	adj [][]int
}

func NewGraph(v int) *Graph {
	adj := make([][]int, 0)
	for i := 0; i < v; i++ {
		adj = append(adj, make([]int, 0))
	}
	return &Graph{v: v, e: 0, adj: adj}
}

func ReadGraph(in io.Reader) *Graph {
	reader := bufio.NewReader(in)
	str, _ := reader.ReadString(' ')
	vertexCnt, _ := strconv.Atoi(strings.Split(str, " ")[0])

	graph := NewGraph(vertexCnt)

	str, _ = reader.ReadString(' ')
	edgeCnt, _ := strconv.Atoi(strings.Split(str, " ")[0])

	for i := 0; i < edgeCnt; i++ {
		str, _ = reader.ReadString(' ')
		v, _ := strconv.Atoi(strings.Split(str, " ")[0])

		str, _ = reader.ReadString(' ')
		w, _ := strconv.Atoi(strings.Split(str, " ")[0])

		graph.addEdge(v, w)
	}

	return graph
}

func (graph *Graph) addEdge(v int, w int) {
	graph.adj[v] = append(graph.adj[v], w)
	graph.adj[w] = append(graph.adj[w], v)
	graph.e++
}

func (graph *Graph) getVertexes() []int {
	vertexes := make([]int, 0, graph.v)
	for i := 0; i < graph.v; i++ {
		vertexes = append(vertexes, i)
	}
	return vertexes
}

func (graph *Graph) GetVertexesCnt() int {
	return graph.v
}
