package email

import (
	"bursa.io/testutils"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("Email Tests", t, func() {

		test_id := testutils.SuffixedId("user")
		email := "admin+" + test_id + "@bursa.io"

		Convey("Subscribe()", func() {
			result := Subscribe(email)
			So(result, ShouldEqual, email)
		})
	})
}
