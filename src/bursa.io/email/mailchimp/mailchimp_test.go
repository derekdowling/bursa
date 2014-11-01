package mailchimp

import (
	"bursa.io/testutils"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("Mailchimp Tests", t, func() {

		test_id := testutils.SuffixedId("user")
		email := "admin+" + test_id + "@bursa.io"

		Convey("SubscribeToChimp()", func() {
			result := SubscribeToChimp(email)
			So(result, ShouldEqual, email)
		})
	})
}
