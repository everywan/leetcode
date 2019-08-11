package ls2json

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_ls2Json(t *testing.T) {
	Convey("passed", t, func() {
		So(ls2Json([]string{}), ShouldEqual, "{}")
		So(ls2Json([]string{""}), ShouldEqual, "{}")
		So(ls2Json([]string{"public/test.md"}), ShouldEqual, "{\"public\":{\"test.md\":\"test.md\"}}")
		So(ls2Json([]string{"public/test.md", "test/", "test.md"}), ShouldEqual, "{\"public\":{\"test.md\":\"test.md\"},\"test\":{},\"test.md\":\"test.md\"}")
	})
}
