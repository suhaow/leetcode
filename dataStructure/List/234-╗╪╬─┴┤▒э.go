package List

/*
leetcode: https://leetcode-cn.com/problems/palindrome-linked-list/
note: https://labuladong.gitee.io/algo/高频面试系列/判断回文链表.html
*/

/*
 判断一个链表是否是回文链表
	回文数：顺序和逆序获取到的值相同
思路：
1. 最简单的办法：将原始链表逆序存入新的链表，对比是否相同
2. 借助二叉树后续遍历的思路
*/

//1. 时间O(N) 空间O(N)
/*
var left *ListNode

func traverse(right *ListNode) bool {
	if nil == right {
		return true
	}
	//后续遍历
	res := traverse(right.Next)

	//比较left 和 right 对应的值
	res = res && (right.Val == left.Val)
	left = left.Next
	return res
}

func isPalindrome(head *ListNode) bool {
	left = head
	return traverse(head)
}
*/

//2. 时间O(N) 空间O(1)
/*
思路： 快慢指针 先找到中间节点  已中间节点为新的头进行反转链表，依次比较头节点 和 反转后的链表
*/
func reverse2(head *ListNode) *ListNode {
	var pre *ListNode
	cur, next := head, head

	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head

	//找到中间节点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	/*
		1->2->2->1->nil   偶数
		1->2->3->2->1->nil  基数
		如果长度为奇数，Slow需要再向前一步
	*/
	if fast != nil {
		slow = slow.Next
	}

	//反转slow后的链表
	left, right := head, slow
	right = reverse2(slow)

	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}

//func main() {
//	// 1->2->3->4->5->nil
//	head := &ListNode{
//		Val: 1,
//	}
//	val := []int{2, 3, 2, 1}
//
//	cur := head
//	for _, v := range val {
//		node := &ListNode{
//			Val: v,
//		}
//		cur.Next = node
//		cur = cur.Next
//	}
//
//	printList(head)
//	println(isPalindrome(head))
//}
