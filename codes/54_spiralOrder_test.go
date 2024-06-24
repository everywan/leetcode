package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_spiralOrder(t *testing.T) {
	Convey("passed", t, func() {
		So(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}), ShouldResemble,
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5})
		So(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}}), ShouldResemble,
			[]int{1, 2, 3, 6, 5, 4})
		So(spiralOrder([][]int{{3}, {2}}), ShouldResemble,
			[]int{3, 2})
		So(spiralOrder([][]int{{6, 9, 7}}), ShouldResemble,
			[]int{6, 9, 7})
	})
}
