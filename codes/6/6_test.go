package covert

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_covert(t *testing.T) {
	Convey("passed", t, func() {
		So(convert("012345678", 3), ShouldEqual, "048135726")
	})
}
