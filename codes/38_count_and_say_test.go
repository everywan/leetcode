package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_countAndSay(t *testing.T) {
	Convey("passed", t, func() {
		So(countAndSay(2), ShouldEqual, "11")
		So(countAndSay(3), ShouldEqual, "21")
		So(countAndSay(4), ShouldEqual, "1211")
		So(countAndSay(5), ShouldEqual, "111221")
	})
}
