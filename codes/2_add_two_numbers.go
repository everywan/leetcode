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
// 如果 l1/l2 符合条件, 是否可以直接返回 l1/l2? 涉及到原值更改的问题
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	start := &ListNode{}
	result := start
	round := 0
	for {
		if l1 == nil && l2 == nil && round == 0 {
			break
		}
		if round > 0 {
			result.Val += round
			round = 0
		}
		if l1 != nil {
			result.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			result.Val += l2.Val
			l2 = l2.Next
		}
		if result.Val > 9 {
			result.Val %= 10
			round = 1
		}
		println(result.Val)
		result.Next = &ListNode{}
		result = result.Next
	}
	return start
}
