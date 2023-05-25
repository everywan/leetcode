package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_uniquePaths(t *testing.T) {
	Convey("passed", t, func() {
		So(uniquePaths(2, 3), ShouldEqual, 3)
		So(uniquePaths(3, 2), ShouldEqual, 3)
		So(uniquePaths(3, 7), ShouldEqual, 28)
		So(uniquePaths(7, 3), ShouldEqual, 28)
	})
}
