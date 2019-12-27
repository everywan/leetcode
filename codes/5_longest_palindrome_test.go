package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_longestPalindrome(t *testing.T) {
	Convey("passed", t, func() {
		So(longestPalindrome("asds"), ShouldEqual, "sds")
	})
}
