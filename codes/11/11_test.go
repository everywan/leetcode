package maxarea

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_maxArea(t *testing.T) {
	Convey("passed", t, func() {
		So(maxArea([]int{1, 2, 3, 6, 3}), ShouldEqual, 6)
	})
}
