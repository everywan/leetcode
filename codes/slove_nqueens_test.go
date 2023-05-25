package codes

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_sloveNQueens(t *testing.T) {
	Convey("passed", t, func() {
		fmt.Println(solveNQueens(4))
		fmt.Println(solveNQueens(5))
	})
}
