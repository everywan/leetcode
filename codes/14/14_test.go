package longestcommonprefix

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_longestCommonPrefix(t *testing.T) {
	Convey("passed", t, func() {
		So(longestCommonPrefix([]string{"asd", "awfxc", "asdwec"}), ShouldEqual, "a")
	})
}
