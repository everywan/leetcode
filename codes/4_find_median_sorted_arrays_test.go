package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_findMedianSortedArrays(t *testing.T) {
	Convey("passed", t, func() {
		So(findMedianSortedArrays([]int{1, 2}, []int{3, 4}), ShouldEqual, 2.5)
		So(findMedianSortedArrays([]int{1}, []int{3, 4}), ShouldEqual, 3)
		So(findMedianSortedArrays([]int{1, 3}, []int{2}), ShouldEqual, 2)
		So(findMedianSortedArrays([]int{}, []int{1}), ShouldEqual, 1)
		So(findMedianSortedArrays([]int{1}, []int{1}), ShouldEqual, 1)
		So(findMedianSortedArrays([]int{2, 3, 4}, []int{1}), ShouldEqual, 2.5)
		So(findMedianSortedArrays([]int{3}, []int{-2, -1}), ShouldEqual, -1)
	})
}
