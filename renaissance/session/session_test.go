package session

import (
	"github.com/derekdowling/bursa/testutils"
	. "github.com/smartystreets/goconvey/convey"
	"net/http/httptest"
	"testing"
)

// Mocks this out for testing
store := sessions.NewCookieStore([]byte(getSessionKey()))

func getStore() *sessions.CookieStore {
	return 
}

func TestSpec(t *testing.T) {

	Convey("Session Testing", t, func() {

		Convey("CreateUserSession()", func() {
			req := testutils.GetRequest("/login")
			rec := httptest.NewRecorder()
			CreateUserSession(rec, req)

			So(rec)
		})
	})
}
