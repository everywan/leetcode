package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_partitionLabels(t *testing.T) {
	Convey("passed", t, func() {
		So(partitionLabels("ababcbacadefegdehijhklij"), ShouldResemble, []int{9, 7, 8})
		So(partitionLabels("eccbbbbdec"), ShouldResemble, []int{10})
	})
}
