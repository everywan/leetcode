package codes

/*
160. 相交链表
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。
如果两个链表不存在相交节点，返回 null 。
*/

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
