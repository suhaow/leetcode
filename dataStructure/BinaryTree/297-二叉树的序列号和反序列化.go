package BinaryTree

import (
	"math"
)

/*
note:https://mp.weixin.qq.com/s/OlpaDhPDTJlQ5MJ8tsARlA
leetcode:https://leetcode-cn.com/problems/maximum-binary-tree/
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
给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：
1. 二叉树的根是数组中的最大元素
2. 左子树是通过数组中最大值左边部分构造出的最大二叉树
3. 右子树是通过数组中最大值右边部分构造出的最大二叉树

[3, 2, 1, 6, 0, 5]
   6
  / \
3     5
 \   /
  2 0
   \
    1
*/
func build(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	//find max val
	maxVal := math.MinInt32
	index := -1
	for i := left; i <= right; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			index = i
		}
	}

	//construct root node
	root := &TreeNode{
		Val: maxVal,
	}

	root.Left = build(nums, left, index-1)
	root.Right = build(nums, index+1, right)
	return root
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}
