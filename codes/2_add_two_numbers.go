package codes

/*
https://leetcode.cn/problems/add-two-numbers/?favorite=2ckc81c
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

2->3->6 + 3->4->6 = 5->7->2->1
632+643 = 1275
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head
	bit := 0
	for {
		val := bit
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		if val >= 10 {
			bit = 1
			val -= 10
		} else {
			bit = 0
		}
		curr.Val = val
		if l1 == nil && l2 == nil {
			break
		}
		curr.Next = &ListNode{}
		curr = curr.Next
	}
	if bit > 0 {
		curr.Next = &ListNode{}
		curr = curr.Next
		curr.Val = bit
	}
	return head
}
