package BinaryTree

/*
note:https://mp.weixin.qq.com/s/ioyqagZLYrvdlZyOMDjrPw
leetcode: https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
*/

/*
目的: 使每个节点的val等于原树中大于或等于 当前节点值的和
思路：左中右 遍历BST能顺序获取节点值，可以通过 右中左的顺序逆序获取节点值，通过外部类加变量 然后赋给每个变量
*/

var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	traverse2(root)
	return root
}
func traverse2(root *TreeNode) {
	if nil == root {
		return
	}

	traverse2(root.Right)
	//累加
	sum += root.Val
	//BST 转换为 累加树
	root.Val = sum
	traverse2(root.Left)
}
