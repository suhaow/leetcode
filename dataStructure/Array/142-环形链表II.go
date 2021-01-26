package Array

/*
leetcode: https://leetcode-cn.com/problems/linked-list-cycle-ii/
note:https://mp.weixin.qq.com/s/yLc7-CZdti8gEMGWhd0JTg
*/

/*
目的：返回链表环的起点
思路：一个指针从相遇点继续向前走，同时另一个指针从head开始走，再次相遇即为环的起点

假设环长为k步, 在二者相遇时，fast 走了2k步， slow 走了 k 步
假设相遇点距离环的入口点为m，
则剩下的环长为k-m，
由于慢指针在判断环时走到相遇点走了k步，因此head到相遇点的步长也为k-m步
综上 当相遇后，一个指针从head走，一个从相遇点走，再次相遇即为入环点
*/

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	//1. 先找到相遇点
	var meet *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			meet = slow
			break
		}
	}

	//若无环，则返回 nil
	if meet == nil {
		return nil
	}

	//此时 slow == fast 都位处相遇点
	newHead := head
	for newHead != slow {
		newHead = newHead.Next
		slow = slow.Next
	}

	return newHead
}
