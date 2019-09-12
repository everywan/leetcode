package reverselistbygroup

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_ls2Json(t *testing.T) {
	Convey("passed", t, func() {
		So(reverseListByGroup([]string{}), ShouldEqual, "{}")
	})
}
