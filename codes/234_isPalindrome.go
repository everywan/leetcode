package codes

import "fmt"

/*
234. 回文链表
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。
如果是，返回 true ；否则，返回 false 。
*/

// tag 解法很多
// 2024-06-24 start: 20:02 end:20:10
// todo 尝试下 快慢指针、递归方法.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	l, r := 0, len(arr)-1
	for l < r {
		if arr[l] != arr[r] {
			return false
		}
		l++
		r--
		fmt.Println(1111, l, r)
	}
	return true
}
