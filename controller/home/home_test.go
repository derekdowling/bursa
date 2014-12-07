package home

import (
	"github.com/derekdowling/bursa/testutils"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("Home Tests", t, func() {

		test_email := testutils.TestEmail("homecontroller")
		test_password := "home_password"

		Convey("marshalForm()", func() {
			form := url.Values{
				"email":    {test_email},
				"password": {test_password},
			}
			req := testutils.FormPostRequest("/signup", form)

			creds := marshalForm(req)
			So(creds, ShouldHaveSameTypeAs, &Credentials{})
			So(creds.Email, ShouldEqual, test_email)
			So(creds.Password, ShouldEqual, test_password)
		})

		Convey("HandlePostSignup()", func() {

			Convey("should work with valid credentials", func() {
				form := url.Values{
					"email":    {test_email},
					"password": {test_password},
				}

				req := testutils.FormPostRequest("/signup", form)
				rec := httptest.NewRecorder()

				HandlePostSignup(rec, req)
				So(rec.Code, ShouldEqual, http.StatusOK)
			})

			Convey("should gracefully handle bad credentials", func() {
				form := url.Values{
					"email":    {"bad_email@blah.com"},
					"password": {""},
				}

				req := testutils.FormPostRequest("/signup", form)
				rec := httptest.NewRecorder()

				HandlePostSignup(rec, req)
				So(rec.Code, ShouldEqual, http.StatusBadRequest)
			})
		})

		Convey("HandlePostLogin", func() {

			Convey("For a successful login", func() {
				form := url.Values{
					"email":    {test_email},
					"password": {test_password},
				}

				req := testutils.FormPostRequest("/login", form)
				rec := httptest.NewRecorder()

				HandlePostLogin(rec, req)
				// TODO: better body checking
				So(rec.Code, ShouldEqual, http.StatusOK)
			})

			Convey("For a bad login", func() {
				form := url.Values{
					"email":    {test_email},
					"password": {"bad_password"},
				}

				req := testutils.FormPostRequest("/login", form)
				rec := httptest.NewRecorder()

				HandlePostLogin(rec, req)
				// TODO: better body checking
				So(rec.Code, ShouldEqual, http.StatusUnauthorized)
			})

		})

	})
}
