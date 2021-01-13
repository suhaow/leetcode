package BinaryTree

/*
leetcode: https://leetcode-cn.com/problems/search-in-a-binary-search-tree/
note:https://mp.weixin.qq.com/s/SuGAvV9zOi4viaeyjWhzDw
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//利用BST特点即可
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}
