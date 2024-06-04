package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_searchMatrix(t *testing.T) {
	Convey("passed", t, func() {
		So(searchMatrix([][]int{
			{1, 2, 3},
			{4, 5, 6},
		}, 5), ShouldEqual, true)
	})
}
