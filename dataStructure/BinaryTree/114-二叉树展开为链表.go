package BinaryTree

/*
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/

/*
note:https://labuladong.gitee.io/algo/数据结构系列/二叉树系列1.html
leetcode:https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
*/

/*
思路:
1. 将 root 左子树和右子树拉平
2. 将 root 右子树拼接到左子树下方，然后将整个左子树作为右子树
*/
func flatten(root *TreeNode) {
	if nil == root {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	//1. 左右子树已被拉平为一条链表
	left := root.Left
	right := root.Right

	//将左子树作为右子树
	root.Left = nil
	root.Right = left

	//3. 将原先的右子树 拼接到 当前右子树的末端
	head := root
	for nil != head.Right {
		head = head.Right
	}
	head.Right = right
}
