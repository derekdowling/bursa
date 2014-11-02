package home

import (
	"bursa.io/testutils"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func urlForm(path string, form url.Values) string {
	url := url.URL{
		Host:     "localhost:8080",
		Path:     path,
		RawQuery: form.Encode(),
	}

	return url.String()
}

func TestSpec(t *testing.T) {

	Convey("Home Tests", t, func() {

		test_email := testutils.TestEmail("homecontroller")

		Convey("HandleSignup()", func() {

			Convey("should work with a valid email", func() {
				form := url.Values{"email": {test_email}}

				rec := httptest.NewRecorder()
				url := urlForm("/signup", form)
				req, err := http.NewRequest("GET", url, nil)

				So(err, ShouldBeNil)

				HandleSignup(rec, req)
				So(rec.Code, ShouldEqual, 200)
				// So(w, ShouldEqual, "14")
			})
		})

	})
}
