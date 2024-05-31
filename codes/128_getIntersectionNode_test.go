package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_getIntersectionNode(t *testing.T) {
	Convey("passed", t, func() {
		{
			common := &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
				},
			}
			headA := &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val:  1,
						Next: common,
					},
				},
			}
			headB := &ListNode{
				Val:  3,
				Next: common,
			}
			So(getIntersectionNode(headA, headB), ShouldEqual, common)
		}
	})
}
