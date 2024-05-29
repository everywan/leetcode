package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_myPow(t *testing.T) {
	Convey("passed", t, func() {
		So(permute([]int{1, 2}), ShouldEqual, [][]int{{1, 2}, {2, 1}})
	})
}
