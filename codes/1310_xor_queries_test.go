package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_xorQueries(t *testing.T) {
	Convey("passed", t, func() {
		So(xorOperation(5, 0), ShouldEqual, 8)
	})
}
