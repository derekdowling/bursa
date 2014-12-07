package models

import (
	"github.com/derekdowling/bursa/testutils"
	"github.com/jinzhu/gorm"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("Models Tests", t, func() {

		Convey("should connect", func() {
			db := Connect()
			So(db, ShouldHaveSameTypeAs, gorm.DB{})
		})

		Convey("User Tests", func() {

			var user *User
			email := testutils.TestEmail("usermodel")
			password := "test123"

			Convey("CreateUser()", func() {
				id := CreateUser(email, password)

				db := Connect()
				db.First(&user, id)

				So(user, ShouldHaveSameTypeAs, &User{})
				So(user.Email, ShouldEqual, email)
				So(user.Password, ShouldNotBeBlank)
			})

			Convey("FindUser()", func() {

				Convey("should find valid user", func() {
					newUser := FindUser(user.Id)
					So(newUser, ShouldHaveSameTypeAs, &User{})
					So(newUser, ShouldResemble, user)
				})

				Convey("on no result", func() {
					user := FindUser(123)
					So(user, ShouldBeNil)
				})
			})

			Convey("FindUserByEmail()", func() {

				Convey("should find valid user", func() {
					newUser := FindUserByEmail(user.Email)
					So(newUser, ShouldHaveSameTypeAs, &User{})
					So(newUser, ShouldResemble, user)
				})

				Convey("on no result", func() {
					user := FindUserByEmail("bademail@bursa.io")
					So(user, ShouldBeNil)
				})
			})

			Convey("FindUserByCreds()", func() {

				Convey("should find valid user", func() {
					newUser := FindUserByCreds(user.Email, password)
					So(newUser, ShouldHaveSameTypeAs, &User{})
					So(newUser, ShouldResemble, user)
				})

				Convey("on no result", func() {
					newUser := FindUserByCreds(user.Email, "badpassword")
					So(newUser, ShouldBeNil)
				})
			})
		})
	})
}
