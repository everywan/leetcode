package codes

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_nextPermutation(t *testing.T) {
	toString := func(nums []int) string {
		return fmt.Sprint(nums)
	}

	Convey("passed", t, func() {
		nums := []int{1, 2, 3}
		//nextPermutation(nums)
		//So(toString(nums), ShouldEqual, "[1 3 2]")

		//nums = []int{3, 2, 1}
		//nextPermutation(nums)
		//So(toString(nums), ShouldEqual, "[1 2 3]")

		//nums = []int{1, 3, 2}
		//nextPermutation(nums)
		//So(toString(nums), ShouldEqual, "[2 1 3]")

		//nums = []int{2, 3, 1}
		//nextPermutation(nums)
		//So(toString(nums), ShouldEqual, "[3 1 2]")

		nums = []int{2, 2, 7, 5, 4, 3, 2, 2, 1}
		nextPermutation(nums)
		So(toString(nums), ShouldEqual, "[2 3 1 2 2 2 4 5 7]")
	})
}
