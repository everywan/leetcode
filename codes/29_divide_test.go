package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_divide(t *testing.T) {
	Convey("passed", t, func() {
		So(divide(10, 3), ShouldEqual, 3)
		So(divide(3, 3), ShouldEqual, 1)
		So(divide(1, 3), ShouldEqual, 0)
		So(divide(3, 1), ShouldEqual, 3)
		So(divide(-3, 3), ShouldEqual, -1)
		So(divide(-3, 1), ShouldEqual, -3)
		// 溢出控制
		So(divide(-1*(1<<31), -1), ShouldEqual, 1<<31-1)
	})
}
