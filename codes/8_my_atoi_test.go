package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_myAtoi(t *testing.T) {
	Convey("passed", t, func() {
		So(myAtoi("1"), ShouldEqual, 1)
		So(myAtoi("-1"), ShouldEqual, -1)
		So(myAtoi("   -1"), ShouldEqual, -1)
		So(myAtoi("123 12a"), ShouldEqual, 123)
		So(myAtoi("123 a"), ShouldEqual, 123)
		So(myAtoi("a 123"), ShouldEqual, 0)
		So(myAtoi("-2147483649"), ShouldEqual, -1*(1<<31))
		So(myAtoi("-91283472332"), ShouldEqual, -2147483648)
		So(myAtoi("2147483800"), ShouldEqual, 2147483647)
	})
}
