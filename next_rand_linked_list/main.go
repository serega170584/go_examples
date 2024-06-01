package main

type Node struct {
	val    int
	next   *Node
	random *Node
}

func main() {
	node3 := &Node{
		val: 3,
	}
	node2 := &Node{
		val:  2,
		next: node3,
	}
	node1 := &Node{
		val:  1,
		next: node2,
	}
	node3.random = node1
	node2.random = node3
	node1.random = node3
}

func copyList(node *Node) *Node {
	var first *Node
	var prev *Node
	var newNode *Node
	links := make(map[*Node]*Node)
	for node != nil {
		if _, ok := links[node]; ok {
			newNode = links[node]
			*newNode = Node{val: node.val}
		} else {
			newNode = &Node{
				val: node.val,
			}
			links[node] = newNode
		}

		if _, ok := links[node.random]; ok {
			newNode.random = links[node.random]
		} else {
			links[node.random] = &Node{}
		}

		if first == nil {
			first = newNode
		}

		if prev != nil {
			prev.next = newNode
		}

		prev = node
		node = node.next
	}

	return first
}
