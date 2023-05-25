package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	Convey("passed", t, func() {
		So(lengthOfLongestSubstring("a"), ShouldEqual, 1)
		So(lengthOfLongestSubstring("aa"), ShouldEqual, 1)
		So(lengthOfLongestSubstring("aba"), ShouldEqual, 2)
		So(lengthOfLongestSubstring("abcabcbb"), ShouldEqual, 3)
	})
}
