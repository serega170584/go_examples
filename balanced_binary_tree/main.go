package main

import (
	"bufio"
	"os"
	"strconv"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())
	tree := make([]*TreeNode, cnt)

	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	curNode := &TreeNode{val: val}

	var curPointerInd, pointerInd int

	tree[0] = curNode
	pointerInd++

	for i := 1; i < cnt; i++ {
		scanner.Scan()
		str := scanner.Text()

		var addNode *TreeNode

		if i%2 == 1 && str != "nil" {
			val, _ = strconv.Atoi(str)
			addNode = &TreeNode{val: val}
			curNode.left = addNode
		}

		if i%2 == 0 && str != "nil" {
			val, _ = strconv.Atoi(str)
			addNode = &TreeNode{val: val}
			curNode.right = addNode
		}

		tree[pointerInd] = addNode
		pointerInd++

		if i%2 == 0 {
			curPointerInd++
			curNode = tree[curPointerInd]
		}
	}
}
