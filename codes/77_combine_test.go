package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_combine(t *testing.T) {
	Convey("passed", t, func() {
		So(combineForLeetcode(4, 1), ShouldResemble,
			[][]int{{1}, {2}, {3}, {4}})
		So(combineForLeetcode(4, 2), ShouldResemble,
			[][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}})
	})
}
