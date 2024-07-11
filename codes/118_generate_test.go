package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_generate(t *testing.T) {
	Convey("passed", t, func() {
		So(generate(5), ShouldResemble, [][]int{
			{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1},
		})
	})
}
