package BinaryTree

/*
leetcode: https://leetcode-cn.com/problems/validate-binary-search-tree/
note:https://mp.weixin.qq.com/s/SuGAvV9zOi4viaeyjWhzDw
*/

func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

func isValid(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if nil == root {
		return true
	}

	// 若 root.val 不符合 max 和 min 的限制，说明不是合法 BST
	if min != nil && root.Val <= min.Val {
		return false
	}

	if max != nil && root.Val >= max.Val {
		return false
	}

	//限定root值作为左子树的最大值 以及 右子树的最小值
	return isValid(root.Left, min, root) &&
		isValid(root.Right, root, max)
}
