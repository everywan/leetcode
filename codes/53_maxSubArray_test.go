package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_maxSubArray(t *testing.T) {
	Convey("passed", t, func() {
		So(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), ShouldEqual, 6)
		So(maxSubArray([]int{-2, 1}), ShouldEqual, 1)
		So(maxSubArray([]int{-1, 0, -2}), ShouldEqual, 0)
	})
}
