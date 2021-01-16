package BinaryTree

import "math"

/*
leetcode: https://leetcode-cn.com/problems/count-complete-tree-nodes/
note:https://mp.weixin.qq.com/s/xW2fbE3v4JhMSKfxoxIHBg
*/

/*
完全二叉树节点个数
1. 计算出左树的高度 和右树的高度
2. 如果左右子树高度相同，则使用满二叉树的计算方式，否则使用普通二叉树的计算方式
*/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHigh, rightHigh := 0, 0
	left, right := root, root
	//求得左右子树高度
	for left != nil {
		left = left.Left
		leftHigh++
	}

	for right != nil {
		right = right.Right
		rightHigh++
	}

	// 满二叉树通过公式计算
	if leftHigh == rightHigh {
		return int(math.Pow(2, float64(leftHigh))) - 1
	}

	//普通二叉树计算
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

/*
普通二叉树节点个数
*/
func countNodesNormal(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodesNormal(root.Left) + countNodesNormal(root.Right)
}

/*
满二叉树
求得树的高度直接计算
*/
func countNodesFull(root *TreeNode) int {
	if root == nil {
		return 0
	}

	high := 0
	for root != nil {
		high++
		root = root.Right
	}
	return int(math.Pow(2, float64(high))) - 1
}
