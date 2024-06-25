package codes

/*
206. 反转链表
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
*/

// tag 需要回顾
// 2024-06-25 start: 21:07 end:
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}
