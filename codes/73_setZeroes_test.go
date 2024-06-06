package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_setZeroes(t *testing.T) {
	Convey("passed", t, func() {
		before := [][]int{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		}
		setZeroes(before)
		So(before, ShouldResemble, [][]int{
			{0, 0, 0, 0},
			{0, 4, 5, 0},
			{0, 3, 1, 0},
		})
	})
}
