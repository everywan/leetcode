package threesum

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_threeSum(t *testing.T) {
	Convey("passed", t, func() {
		expected := [][]int{[]int{0, 0, 0}, []int{-5, 4, 1}, []int{-4, 4, 0}, []int{-4, 3, 1}, []int{-2, 4, -2}, []int{-2, 1, 1}}
		So(threeSum([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}), ShouldResemble, expected)
	})
}
