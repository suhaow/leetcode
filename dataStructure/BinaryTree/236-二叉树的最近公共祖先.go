package BinaryTree

/*
note: https://mp.weixin.qq.com/s/9RKzBcr3I592spAsuMH45g
leetcode: https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
*/

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}

	if p == root || q == root {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	//situation1: p q都在root的子树中 由于是后续遍历 所以root一定是p q的最近公共组件
	if left != nil && right != nil {
		return root
	}

	//situation2: p q都不在root的子树中，返回Nil
	if left == nil && right == nil {
		return nil
	}

	//situation3: p q只有一个存在root为根的子树中，返回该节点
	if left == nil {
		return right
	}
	return left
}
