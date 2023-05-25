package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_decode(t *testing.T) {
	Convey("passed", t, func() {
		So(decode([]int{1, 2, 3}, 1), ShouldResemble, []int{1, 0, 2, 1})
		So(decode([]int{6, 2, 7, 3}, 4), ShouldResemble, []int{4, 2, 0, 7, 4})
	})
}
