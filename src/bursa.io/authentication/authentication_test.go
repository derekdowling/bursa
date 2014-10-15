package authentication

import (
	"code.google.com/p/go.crypto/bcrypt"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"strings"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("Authentication Testing", t, func() {

		Convey("generateSalt()", func() {
			salt := generateSalt()
			So(salt, ShouldNotBeBlank)
			So(len(salt), ShouldEqual, SaltLength)
		})

		Convey("combine()", func() {
			salt := generateSalt()
			password := "boomchuckalucka"
			expectedLength := len(salt) + len(password)
			combo := combine(salt, password)

			So(combo, ShouldNotBeBlank)
			So(len(combo), ShouldEqual, expectedLength)
			So(strings.HasPrefix(combo, salt), ShouldBeTrue)
		})

		Convey("hashPassword()", func() {
			combo := combine(generateSalt(), "hershmahgersh")
			hash := hashPassword(combo)
			So(hash, ShouldNotBeBlank)

			cost, err := bcrypt.Cost([]byte(hash))
			if err != nil {
				log.Print(err)
			}
			So(cost, ShouldEqual, EncryptCost)
		})

		Convey("CreatePassword()", func() {
			passString := "mmmPassword1"
			password := CreatePassword(passString)
			pass_struct := new(Password)

			So(password, ShouldHaveSameTypeAs, pass_struct)
			So(password.hash, ShouldNotBeBlank)
			So(password.salt, ShouldNotBeBlank)
			So(len(password.salt), ShouldEqual, SaltLength)
		})

		Convey("PasswordMatch()", func() {
			password := "megaman49"
			passwordMeta := CreatePassword(password)

			So(PasswordMatch(password, passwordMeta), ShouldBeTrue)
			So(PasswordMatch("lolfail", passwordMeta), ShouldBeFalse)
			So(PasswordMatch("Megaman49", passwordMeta), ShouldBeFalse)
		})
	})
}
