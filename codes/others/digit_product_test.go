package others

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_digitProduct(t *testing.T) {
	Convey("passed", t, func() {
		So(digitProduct(36), ShouldEqual, 49)
		So(digitProduct(100), ShouldEqual, 455)
		So(digitProduct(17), ShouldEqual, -1)
	})
}
