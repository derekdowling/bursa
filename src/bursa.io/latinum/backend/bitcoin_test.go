package backend

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

		Convey("ProvisionWallet", func() {
			Convey("Should create a new user account", func() {

			})

			Convey("Should create a new user account", func() {

			})

			Convey("Should create a new user account", func() {

			})
		})
	})
}
