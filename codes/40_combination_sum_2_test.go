package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_combinationSum2(t *testing.T) {
	Convey("passed", t, func() {
		So(combinationSum2([]int{}, 0), ShouldResemble, [][]int{})
		So(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8), ShouldResemble,
			[][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}})
		So(combinationSum2([]int{3, 1, 3, 5, 1, 1}, 8), ShouldResemble,
			[][]int{{1, 1, 1, 5}, {1, 1, 3, 3}, {3, 5}})
	})
}
