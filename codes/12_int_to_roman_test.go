package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_intToRoman(t *testing.T) {
	Convey("passed", t, func() {
		So(intToRoman(1), ShouldEqual, RomanI)
		So(intToRoman(9), ShouldEqual, "IX")
		So(intToRoman(18), ShouldEqual, "XVIII")
		So(intToRoman(19), ShouldEqual, "IXX")
		SkipSo(intToRoman(1994), ShouldEqual, "MCMXCIV")
	})
}
