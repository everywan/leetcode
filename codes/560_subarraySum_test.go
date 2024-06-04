package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_subarraySum(t *testing.T) {
	Convey("passed", t, func() {
		So(subarraySum([]int{1, 1, 1}, 2), ShouldEqual, 2)
		So(subarraySum([]int{1, 2, 3}, 3), ShouldEqual, 2)
		So(subarraySum([]int{1, -1, 0}, 0), ShouldEqual, 3)
	})
}
