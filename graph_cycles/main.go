package main

type Graph struct {
	vertexesCnt   int
	vertexesList  [][]int
	marked        []bool
	coloured      []bool
	markedInCycle []bool
	cycles        [][]int
}

func NewGraph(vertexesCnt int, edgesCnt int, edges [][2]int) *Graph {
	vertexesList := make([][]int, vertexesCnt)
	marked := make([]bool, vertexesCnt)
	markedInCycle := make([]bool, vertexesCnt)
	coloured := make([]bool, vertexesCnt)
	for i := 0; i < edgesCnt; i++ {
		firstVertex := edges[i][0]
		secondVertex := edges[i][1]
		vertexesList[firstVertex] = append(vertexesList[firstVertex], secondVertex)
	}
	return &Graph{
		vertexesCnt:   vertexesCnt,
		vertexesList:  vertexesList,
		marked:        marked,
		coloured:      coloured,
		markedInCycle: markedInCycle,
	}
}

func (g *Graph) dfs(i int) {
	list := g.vertexesList[i]
	g.marked[i] = true
	colouredCnt := 0
	for _, v := range list {
		if !g.marked[v] {
			g.dfs(v)
		}

		if g.coloured[v] {
			colouredCnt++
		}
	}

	if len(list) == colouredCnt {
		g.coloured[i] = true
	}
}

func (g *Graph) visitCycle(root int, i int, cycle []int) {
	if i == root {
		g.cycles = append(g.cycles, cycle)
	}

	if i == -1 {
		i = root
	}

	cycle = append(cycle, i)

	g.markedInCycle[i] = true
	list := g.vertexesList[i]
	for _, v := range list {
		if !g.coloured[v] {
			newCycle := make([]int, len(cycle))
			copy(newCycle, cycle)
			g.visitCycle(root, v, newCycle)
		}
	}
}

func (g *Graph) findCycles() [][]int {
	for i := 0; i < g.vertexesCnt; i++ {
		if !g.coloured[i] && !g.markedInCycle[i] {
			cycle := make([]int, 0)
			g.visitCycle(i, i, cycle)
		}
	}
	return g.cycles
}

func main() {

}
