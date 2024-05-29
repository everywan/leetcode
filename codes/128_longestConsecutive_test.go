package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_longestConsecutive(t *testing.T) {
	Convey("passed", t, func() {
		// So(longestConsecutive([]int{100, 4, 200, 1, 3, 2}), ShouldEqual, 4)
		// So(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}), ShouldEqual, 9)
		So(longestConsecutive([]int{4, 0, -4, -2, 2, 5, 2, 0, -8, -8, -8, -8, -1, 7, 4, 5, 5, -4, 6, 6, -3}),
			ShouldEqual, 5)
	})
}
