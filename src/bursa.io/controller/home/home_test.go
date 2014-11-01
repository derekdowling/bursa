package home

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("Config Tests", t, func() {

		Convey("HandleSignup()", func() {

			test_email := "test"

			query := url.Values{"email": {test_email}}
			req, err := http.NewRequest("http://localhost:8080/signup", query)
			if err != nil {
				log.Fatal(err)
			}

			w := httptest.NewRecorder()
			HandleSignup(w, req)
		})

	})
}
