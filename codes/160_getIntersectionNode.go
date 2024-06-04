package codes

/*
160. 相交链表
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。
如果两个链表不存在相交节点，返回 null 。
*/

// tag: 非最佳解法
// 2024-05-31 start: 18:59 end:19:18
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	set := map[*ListNode]struct{}{}
	for headA != nil || headB != nil {
		if headA != nil {
			if _, ok := set[headA]; ok {
				return headA
			}
			set[headA] = struct{}{}
			headA = headA.Next
		}
		if headB != nil {
			if _, ok := set[headB]; ok {
				return headB
			}
			set[headB] = struct{}{}
			headB = headB.Next
		}
	}
	return nil
}

// Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// 最佳思路解法
// func getIntersectionNode2(headA, headB *ListNode) *ListNode {
// 	firstA, firstB := headA, headB
// 	for {
// 		if headA == headB {
// 			return headA
// 		}
// 		if headA != nil {
// 			headA = headA.Next
// 		} else {
// 			headA = firstB
// 		}
// 		if headB != nil {
// 			headB = headB.Next
// 		} else {
// 			headB = firstA
// 		}
// 	}
// }
