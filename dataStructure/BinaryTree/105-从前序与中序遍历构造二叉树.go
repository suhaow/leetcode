package BinaryTree

/*
note:https://mp.weixin.qq.com/s/OlpaDhPDTJlQ5MJ8tsARlA
leetcode:https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
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
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]

返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build1(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build1(preorder []int, pStart, pEnd int, inorder []int, iStart, iEnd int) *TreeNode {
	if pStart > pEnd {
		return nil
	}
	//root 就是前序中的第一个元素
	rootVal := preorder[pStart]
	root := &TreeNode{
		Val: rootVal,
	}

	//找到root 在中序遍历中的下标，从而区分左右子树
	rootIndex := 0
	for i := iStart; i <= iEnd; i++ {
		if inorder[i] == rootVal {
			rootIndex = i
		}
	}

	//构造左右子树
	root.Left = build1(preorder, pStart+1, pStart+rootIndex-iStart, inorder, iStart, rootIndex-1)
	root.Right = build1(preorder, pStart+rootIndex-iStart+1, pEnd, inorder, rootIndex+1, iEnd)
	return root
}
