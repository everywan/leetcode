package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_combinationSum(t *testing.T) {
	Convey("passed", t, func() {
		So(combinationSum([]int{}, 0), ShouldResemble, [][]int{})
		So(combinationSum([]int{1, 2, 3}, 3), ShouldResemble, [][]int{
			[]int{1, 1, 1}, []int{1, 2}, []int{3}})
		So(combinationSum([]int{3, 2, 1}, 2), ShouldResemble, [][]int{
			[]int{1, 1}, []int{2}})
		// 用 slice append 时, 切片值被更改
		So(combinationSum([]int{7, 3, 2}, 18), ShouldResemble, [][]int{
			[]int{2, 2, 2, 2, 2, 2, 2, 2, 2}, []int{2, 2, 2, 2, 2, 2, 3, 3},
			[]int{2, 2, 2, 2, 3, 7}, []int{2, 2, 2, 3, 3, 3, 3}, []int{2, 2, 7, 7},
			[]int{2, 3, 3, 3, 7}, []int{3, 3, 3, 3, 3, 3}})
	})
}
