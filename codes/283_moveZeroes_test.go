package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_moveZeroes(t *testing.T) {
	Convey("passed", t, func() {
		{
			nums := []int{0, 1, 0, 3, 12}
			moveZeroes(nums)
			So(nums, ShouldResemble, []int{1, 3, 12, 0, 0})
		}
	})
}
