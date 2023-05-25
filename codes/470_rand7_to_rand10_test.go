package codes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_rand10(t *testing.T) {
	Convey("passed", t, func() {
		for i := 0; i < 20; i++ {
			t.Log(rand10())
		}
	})
}
