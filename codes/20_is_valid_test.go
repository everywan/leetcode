package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_isValid(t *testing.T) {
	Convey("passed", t, func() {
		So(isValid(""), ShouldEqual, true)
		So(isValid("[(){([])}]"), ShouldEqual, true)
		So(isValid("[(){([)]}]"), ShouldEqual, false)
		So(isValid("(){}"), ShouldEqual, true)
		So(isValid("}"), ShouldEqual, false)
		So(isValid("["), ShouldEqual, false)
	})
}
