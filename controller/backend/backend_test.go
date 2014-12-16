package backend

import (
	"github.com/derekdowling/bursa/testutils"
	. "github.com/smartystreets/goconvey/convey"
	"net/http/httptest"
	"net/url"
	"testing"
	"net/http"
)

func TestSpec(t *testing.T) {

	Convey("Home Tests", t, func() {

		test_email := testutils.TestEmail("homecontroller")

		Convey("HandleCreate()", func() {

			Convey("should save a wallet", func() {
				form := url.Values{"email": {test_email}}

				req, err := testutils.FormPostRequest("/signup", form)
				rec := httptest.NewRecorder()

				So(err, ShouldBeNil)
				HandleSignup(rec, req)
				So(rec.Code, ShouldEqual, http.StatusOK)
			})
		})
	})
}
