package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if nil == root {
		return nil
	}
	connectTwoNode(root.Left, root.Right)
	return root
}

func connectTwoNode(node1 *Node, node2 *Node) {
	if nil == node1 || nil == node2 {
		return
	}

	node1.Next = node2
	//连接相同父节点的两个子节点
	connectTwoNode(node1.Left, node1.Right)
	connectTwoNode(node2.Left, node2.Right)

	//连接跨越父节点的两个子节点
	connectTwoNode(node1.Right, node2.Left)
}
