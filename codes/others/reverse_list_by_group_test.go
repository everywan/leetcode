package others

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_reverseListByGroup(t *testing.T) {
	Convey("passed", t, func() {
		So(reverseListByGroup([]string{}), ShouldEqual, "")
	})
}
