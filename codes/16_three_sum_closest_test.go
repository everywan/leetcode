package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_threeSumClosest(t *testing.T) {
	Convey("passed", t, func() {
		So(threeSumClosest([]int{-1, 2, 1, -4}, 1), ShouldEqual, 2)
	})
}
