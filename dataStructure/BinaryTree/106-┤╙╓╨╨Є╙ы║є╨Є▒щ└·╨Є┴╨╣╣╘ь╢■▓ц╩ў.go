package BinaryTree

/*
note:https://mp.weixin.qq.com/s/OlpaDhPDTJlQ5MJ8tsARlA
leetcode:https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
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
中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]

    3
   / \
  9  20
    /  \
   15   7

*/

func buildTree1(inorder []int, postorder []int) *TreeNode {
	return build2(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func build2(inorder []int, iStart, iEnd int, postorder []int, pStart, pEnd int) *TreeNode {
	if iStart > iEnd {
		return nil
	}

	//后序遍历中最后一个节点即为当前的根节点
	rootVal := postorder[pEnd]

	//找到中序遍历中根节点的下标，区分左右子树
	rootIndex := 0
	for i := iStart; i <= iEnd; i++ {
		if inorder[i] == rootVal {
			rootIndex = i
			break
		}
	}

	root := &TreeNode{
		Val: rootVal,
	}

	//左子树的节点个数
	leftSize := rootIndex - iStart

	//递归左右子树
	root.Left = build2(inorder, iStart, rootIndex-1, postorder, pStart, pStart+leftSize-1)
	root.Right = build2(inorder, rootIndex+1, iEnd, postorder, pStart+leftSize, pEnd-1)
	return root
}
