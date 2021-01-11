package List

//https://labuladong.gitee.io/algo/数据结构系列/递归反转链表的一部分.html

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	ptr := head
	for ptr != nil {
		println(ptr.Val)
		ptr = ptr.Next
	}
}

/*
一、递归反转整个链表
*/
func reverseAll(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	node := reverseAll(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

/*
二、反转链表前 N 个节点
*/
var successor *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		// 记录第 n+1 个节点
		successor = head.Next
		return head
	}

	last := reverseN(head.Next, n-1)

	head.Next.Next = head
	// 让反转之后的 head 节点和后面的节点连起来
	head.Next = successor

	return last
}

/*
三、反转链表的一部分
*/
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// 找到需要反转的头节点后，当作逆转链表的前N个节点问题处理
	if m == 1 {
		return reverseN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

//func main() {
//	// 1->2->3->4->5->nil
//	head := &ListNode{
//		Val: 1,
//	}
//	cur := head
//	for i := 2; i <= 5; i++ {
//		node := &ListNode{
//			Val: i,
//		}
//		cur.Next = node
//		cur = cur.Next
//	}
//
//	// 5->4->3->2->1->nil
//	//newHead := reverseAll(head)
//
//	// 3->2->1->4->5
//	//newHead := reverseN(head, 3)
//
//	newHead := reverseBetween(head, 2, 3)
//	printList(newHead)
//}
