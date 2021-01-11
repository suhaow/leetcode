package BinaryTree

import (
	"strconv"
)

/*
note: https://mp.weixin.qq.com/s/LJbpo49qppIeRs-FbgjsSQ
leetcode: https://leetcode-cn.com/problems/find-duplicate-subtrees/
*/

/*
思路:
1. 利用外部变量存储每个子树序列化的结果
2. 对于每个节点 判断自己是否与序列化的结果重复
*/

/*
       1
      / \
     2   3
    /   / \
   4   2   4
      /
     4

*/

//leetcode 中该题不可使用全局变量，否则多个case中会受影响
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var res []*TreeNode
	countMap := make(map[string]int)
	traverse(root, &res, countMap)
	return res
}

func traverse(root *TreeNode, res *[]*TreeNode, countMap map[string]int) string {
	if nil == root {
		return "#"
	}

	left := traverse(root.Left, res, countMap)
	right := traverse(root.Right, res, countMap)

	subTree := left + "," + right + "," + strconv.Itoa(root.Val)
	if count, ok := countMap[subTree]; !ok {
		countMap[subTree] = 1
	} else {
		if 1 == count {
			*res = append(*res, root)
		}
		countMap[subTree]++
	}
	//出现次数++
	return subTree
}
