package removenthfromend

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_removeNthFromEnd(t *testing.T) {
	SkipConvey("passed", t, func() {
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
	})
}
