package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := strings.NewReader(os.Args[1])
	s, _ := strconv.Atoi(os.Args[2])
	graph := ReadGraph(in)
	search := NewSearch(g, s)
	for v := 0; v < graph.GetVertexesCnt(); v++ {
		if search.marked(v) {
			fmt.Println(v + " ")
		}
	}
	if search.Count() != graph.GetVertexesCnt() {
		fmt.Println("НЕ")
	}
	fmt.Println("связный")
}
