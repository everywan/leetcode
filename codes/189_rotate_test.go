package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_rotate(t *testing.T) {
	Convey("passed", t, func() {
		nums := []int{1, 2, 3, 4, 5, 6, 7}
		rotate(nums, 3)
		So(nums, ShouldResemble, []int{5, 6, 7, 1, 2, 3, 4})
	})
}
