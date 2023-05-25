package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_xorOperation(t *testing.T) {
	Convey("passed", t, func() {
		So(xorOperation(5, 0), ShouldEqual, 8)
	})
}
