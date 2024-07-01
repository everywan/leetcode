package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_isPalindrome(t *testing.T) {
	Convey("passed", t, func() {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		}
		So(isPalindrome(head), ShouldEqual, false)
	})
}
