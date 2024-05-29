package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_groupAnagrams(t *testing.T) {
	Convey("passed", t, func() {
		So(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}), ShouldResemble,
			[][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}})
	})
}
