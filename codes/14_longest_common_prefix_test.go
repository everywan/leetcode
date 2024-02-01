package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_longestCommonPrefix(t *testing.T) {
	Convey("passed", t, func() {
		So(longestCommonPrefix([]string{}), ShouldEqual, "")
		So(longestCommonPrefix([]string{"a", "ab", "abc"}), ShouldEqual, "a")
	})
}
