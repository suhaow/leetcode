package BinaryTree

/*
note:https://mp.weixin.qq.com/s/ioyqagZLYrvdlZyOMDjrPw
leetcode:https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
*/

/*
题意：给定二叉搜索树  找到第k小的元素
*/

func kthSmallest(root *TreeNode, k int) int {
	traverse1(root, k)
	return res
}

// 结果
var res int

// 当前元素排名
var rank int

func traverse1(root *TreeNode, k int) {
	if nil == root {
		return
	}

	//找到最小节点
	traverse1(root.Left, k)
	rank++
	if k == rank {
		res = root.Val
		return
	}
	traverse1(root.Right, k)
}
