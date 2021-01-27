package Array

/*
leetcode: https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
note:https://mp.weixin.qq.com/s/yLc7-CZdti8gEMGWhd0JTg
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head

	//快指针先走n步
	for fast != nil && n > 0 {
		fast = fast.Next
		n--
	}

	//避免 n > 链表长度
	if n > 0 {
		return nil
	}
	//删除头节点情况
	if fast == nil {
		head = slow.Next
		return head
	}

	//slow fast 一起走
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}
