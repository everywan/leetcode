package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_merge(t *testing.T) {
	Convey("passed", t, func() {
		So(merge([][]int{
			{1, 3},
			{2, 6},
			{8, 10},
		}), ShouldResemble, [][]int{{1, 6}, {8, 10}})

		So(merge([][]int{
			{1, 4},
			{0, 4},
		}), ShouldResemble, [][]int{{0, 4}})
	})
}
