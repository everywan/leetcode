package searchinsert

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_searchInsert(t *testing.T) {
	Convey("passed", t, func() {
		So(searchInsert([]int{1, 2, 3, 6}, 2), ShouldEqual, 1)
		So(searchInsert([]int{}, 2), ShouldEqual, 0)
		So(searchInsert([]int{1, 2, 3}, 4), ShouldEqual, 3)
	})
}
