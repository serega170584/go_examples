package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

//          1
//     2            3
// 4      5      6     nil
//7  8 nil nil nil nil
//9 10

// 1 2 3 4 5 6 nil 7 8 nil nil nil nil 9 10

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

			for curNode == nil {
				curPointerInd++
				curNode = tree[curPointerInd]
			}
		}
	}

	fmt.Println(isBalanced(tree[0]))
}

func isBalanced(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	isBalancedNode, leftLen := isBalanced(root.left)
	if !isBalancedNode {
		return false, 0
	}

	isBalancedNode, rightLen := isBalanced(root.right)
	if !isBalancedNode {
		return false, 0
	}

	if leftLen-rightLen < -1 || 1 < leftLen-rightLen {
		return false, 0
	}

	max := leftLen
	if rightLen > max {
		max = rightLen
	}

	return true, max + 1
}
