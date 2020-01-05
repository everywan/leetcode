package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_findSubstring(t *testing.T) {
	Convey("passed", t, func() {
		//So(findSubstring("arfoothefoobarman", []string{"foo", "bar"}), ShouldResemble, []int{8})
		//So(findSubstring("barfoothefoobarman", []string{"foo", "bar"}), ShouldResemble, []int{0, 9})
		So(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}),
			ShouldResemble, []int{8})
	})
}
