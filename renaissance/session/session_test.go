package session

import (
	"github.com/derekdowling/bursa/models"
	"github.com/derekdowling/bursa/testutils"
	"github.com/gorilla/sessions"
	. "github.com/smartystreets/goconvey/convey"
	"net/http/httptest"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("Session Testing", t, func() {

		store := sessions.NewCookieStore([]byte(getSessionKey()))
		session := New()
		session.Store = store
		email := testutils.TestEmail("session_test")
		password := "testpass"

		user_id := models.CreateUser(email, password)

		Convey("CreateUserSession()", func() {
			req := testutils.GetRequest("/login")
			rec := httptest.NewRecorder()
			session.CreateUserSession(rec, req, user_id)
			session, err := session.Store.Get(req, "user_session")

			So(err, ShouldBeNil)
			So(session, ShouldHaveSameTypeAs, &sessions.Session{})
		})
	})
}
