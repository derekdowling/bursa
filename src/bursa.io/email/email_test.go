package email

import (
	"bursa.io/testutils"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("Email Tests", t, func() {

		email := testutils.TestEmail("user")

		Convey("Subscribe()", func() {
			result := Subscribe(email)
			So(result, ShouldBeTrue)
		})
	})
}
