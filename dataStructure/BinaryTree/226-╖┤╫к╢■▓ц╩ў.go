package BinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
note:https://labuladong.gitee.io/algo/数据结构系列/二叉树系列1.html
leetcode:https://leetcode-cn.com/problems/invert-binary-tree/
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	/*
		前序遍历，依次交换左右子节点
	*/
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
