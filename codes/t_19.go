package codes

import "fmt"

// 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
// 尝试使用一趟扫描实现
func RemoveNthFromEnd() {
	h := &ListNode{Val: 1, Next: nil}
	h.Next = &ListNode{Val: 2, Next: nil}
	h.Next.Next = &ListNode{Val: 3, Next: nil}
	h.Next.Next.Next = &ListNode{Val: 4, Next: nil}
	h.Next.Next.Next.Next = &ListNode{Val: 5, Next: nil}
	t := removeNthFromEnd(h, 2)
	for t != nil {
		fmt.Println(t.Val)
		t = t.Next
	}
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 简单而言, 只需要两次遍历即可, 一次求长度, 一次删数据.
// 只需一次遍历的方法: 设置 head, headn 两个节点, headn 为第N个节点, 当 headn == tail 时, head就是要寻找的节点.
// 链表操作常用思想, 我之前也遇到过, 包括判断是否是环, 是否有环等,  竟然给忘了.. (判断 head,head2n 是否会相遇).
// 比较蛋疼的是, 我比较容易在下标/边界条件中搞混... 难受
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	headn, result := head, head
	for n > 0 {
		n--
		if headn.Next == nil {
			return head.Next
		}
		headn = headn.Next
	}
	for headn.Next != nil {
		headn = headn.Next
		head = head.Next
	}
	head.Next = head.Next.Next
	return result
}
