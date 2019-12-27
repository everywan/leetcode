package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_reverse(t *testing.T) {
	Convey("passed", t, func() {
		So(reverse(123), ShouldEqual, 321)
	})
}
