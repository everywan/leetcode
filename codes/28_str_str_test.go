package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_strStr(t *testing.T) {
	Convey("passed", t, func() {
		So(strStr("asdas", "1"), ShouldEqual, -1)
		So(strStr("asdas", "asd"), ShouldEqual, 0)
		So(strStr("123asddd1112344111", "dd"), ShouldEqual, 5)
		So(strStr("", ""), ShouldEqual, 0)
		So(strStr("", "1"), ShouldEqual, -1)
		So(strStr("1", "1"), ShouldEqual, 0)
	})
}
