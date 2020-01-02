package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_generateParenthesis(t *testing.T) {
	Convey("passed", t, func() {
		So(generateParenthesis(1), ShouldResemble, []string{"()"})
		So(generateParenthesis(2), ShouldResemble, []string{"(())", "()()"})
		So(generateParenthesis(3), ShouldResemble, []string{"((()))", "(()())",
			"(())()", "()(())", "()()()"})
	})
}
