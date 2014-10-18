package bitcoin

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {
	Convey("Latinum Tests", t, func() {
		Convey("NewLatinum()", func() {
			Convey("Should create a new Latinum instance", func() {
				So(NewLatinum(), ShouldHaveSameTypeAs, &Latinum{})
			})
		})
	})
}
