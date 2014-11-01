package models

import (
	"bursa.io/testutils"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("User Tests", t, func() {

		test_id := testutils.EmailSuffixedId("user")
		email := "admin+" + test_id + "@bursa.io"

		Convey("SubscribeToMail()", func() {
			result := SubscribeToMail(email)
			So(result.Email, ShouldEqual, email)
		})
	})
}
