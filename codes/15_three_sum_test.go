package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_threeSum(t *testing.T) {
	Convey("passed", t, func() {
		So(threeSum([]int{1, 1, -2}), ShouldResemble, []int{1, 1, -2})
	})
}
