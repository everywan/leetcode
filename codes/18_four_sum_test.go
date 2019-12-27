package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_fourSum(t *testing.T) {
	Convey("passed", t, func() {
		So(fourSum([]int{1, 0, -1, 0, -2, 2}, 0), ShouldResemble,
			[][]int{[]int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}})
		So(fourSum([]int{0, 0, 0, 0}, 1), ShouldResemble, [][]int{})
		So(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0), ShouldResemble,
			[][]int{[]int{-3, -2, 2, 3}, []int{-3, -1, 1, 3}, []int{-3, 0, 0, 3}, []int{-3, 0, 1, 2},
				[]int{-2, -1, 0, 3}, []int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}})
		So(fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11), ShouldResemble, [][]int{[]int{-5, -4, -3, 1}})
		So(fourSum([]int{1, 0, -1, 0, -2, 2}, 0), ShouldResemble,
			[][]int{[]int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}})
	})
}
