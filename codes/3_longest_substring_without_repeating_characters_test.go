package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	Convey("passed", t, func() {
		So(lengthOfLongestSubstringV2("a"), ShouldEqual, 1)
		So(lengthOfLongestSubstringV2("aa"), ShouldEqual, 1)
		So(lengthOfLongestSubstringV2("aba"), ShouldEqual, 2)
		So(lengthOfLongestSubstringV2("abcabcbb"), ShouldEqual, 3)
		So(lengthOfLongestSubstringV2("au"), ShouldEqual, 2)
		So(lengthOfLongestSubstringV2("aab"), ShouldEqual, 2)
		So(lengthOfLongestSubstringV2("dvdf"), ShouldEqual, 3)
		So(lengthOfLongestSubstringV2("pwwkew"), ShouldEqual, 3)
	})
}
