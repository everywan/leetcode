package removeDuplicates

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_removeDuplicates(t *testing.T) {
	Convey("passed", t, func() {
		So(removeDuplicates([]int{1, 2, 3, 4}), ShouldEqual, 4)
		So(removeDuplicates([]int{1, 1, 2, 3, 4}), ShouldEqual, 4)
		So(removeDuplicates([]int{1, 1, 2, 2, 3, 4}), ShouldEqual, 4)
	})
}
