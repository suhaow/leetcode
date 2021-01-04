package main

/*
leetcode: https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
note: https://labuladong.gitee.io/algo/高频面试系列/k个一组反转链表.html
*/

/*
	反转链表
*/
func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur, next := head, head

	for cur != nil {
		//记录cur  next
		next = cur.Next

		//反转cur next 指向
		cur.Next = pre

		//更新指针位置
		pre = cur
		cur = next
	}
	return pre
}

/*
	反转 [a, b) 元素
*/
func reverse1(left *ListNode, right *ListNode) *ListNode {
	var pre *ListNode
	cur, next := left, left

	for cur != right {
		//记录cur  next
		next = cur.Next

		//反转cur next 指向
		cur.Next = pre

		//更新指针位置
		pre = cur
		cur = next
	}
	//反转后的头节点
	return pre
}

func reverseGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	left, right := head, head

	for i := 0; i < k; i++ {
		if right == nil {
			return head
		}
		right = right.Next
	}
	//反转前k个
	newHead := reverse1(left, right)
	// 合并结果
	left.Next = reverseGroup(right, k)
	return newHead
}

func main() {
	// 1->2->3->4->5->nil
	head := &ListNode{
		Val: 1,
	}
	cur := head
	for i := 2; i <= 5; i++ {
		node := &ListNode{
			Val: i,
		}
		cur.Next = node
		cur = cur.Next
	}

	// 5->4->3->2->1->nil
	//newHead := reverse(head)

	// 2->1->4->3->5
	newHead := reverseGroup(head, 2)
	printList(newHead)
}
