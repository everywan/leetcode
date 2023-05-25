package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_addTwoNumbers(t *testing.T) {
	Convey("passed", t, func() {
		l1 := &ListNode{3, &ListNode{4, &ListNode{6, nil}}}
		l2 := &ListNode{2, &ListNode{3, &ListNode{6, nil}}}
		r1 := addTwoNumbers(l1, l2)
		So(r1.Val, ShouldEqual, 5)
		So(r1.Next.Val, ShouldEqual, 7)
		So(r1.Next.Next.Val, ShouldEqual, 2)
		So(r1.Next.Next.Next.Val, ShouldEqual, 1)
	})
}
