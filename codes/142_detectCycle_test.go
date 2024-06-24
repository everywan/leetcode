package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_detectCycle(t *testing.T) {
	Convey("passed", t, func() {
		{
			common := &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: -4,
					},
				},
			}
			common.Next.Next.Next = common
			head := &ListNode{
				Val:  3,
				Next: common,
			}
			So(detectCycle(head), ShouldEqual, common)
		}
	})
}
