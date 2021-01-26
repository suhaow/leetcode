package Array

/*
leetcode: https://leetcode-cn.com/problems/linked-list-cycle/
note: https://mp.weixin.qq.com/s/yLc7-CZdti8gEMGWhd0JTg
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
目的：判断是否有环
快慢指针
慢指针一次走一步 快指针一次走两步 若有环二者比相遇
*/
func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}
