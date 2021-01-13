package BinaryTree

/*
leetcode: https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/
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

/*
思路：按照BST特性找到val应该所在的位置插入即可
*/
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if nil == root {
		return &TreeNode{
			Val: val,
		}
	}

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}

	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}
