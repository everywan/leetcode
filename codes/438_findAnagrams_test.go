package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_findAnagrams(t *testing.T) {
	Convey("passed", t, func() {
		So(findAnagrams("cbaebabacd", "abc"), ShouldResemble, []int{0, 6})
		So(findAnagrams("baa", "aa"), ShouldResemble, []int{1})
	})
}
