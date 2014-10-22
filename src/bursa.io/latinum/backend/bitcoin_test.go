package backend

import (
	"log"
	"bursa.io/latinum/vault"
	"bursa.io/models"
	"bursa.io/testutils"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {
	// TODO Doing this all over is ugly.
	db, err := models.Connect()
	if err != nil {
		log.Fatalf("Couldn't connect to database during testing", err)
	}

	user_a := models.User{
		Name: testutils.SuffixedId("bitcoin_test_user_a"),
	}
	db.Save(&user_a)

	user_b := models.User{
		Name: testutils.SuffixedId("bitcoin_test_user_a"),
	}
	db.Save(&user_b)

	Convey("Latinum Tests", t, func() {
		Convey("NewLatinum()", func() {
			Convey("Should create a new Latinum instance", func() {
				So(NewLatinum(), ShouldHaveSameTypeAs, &Latinum{})
			})
		})

		Convey("Generate()", func() {
			Convey("Should generate bitcoins for testing", func() {
				// Generate()
			})
		})

		Convey("Send()", func() {
			Convey("Should send bitcoins from a to b", func() {
				// address_a = vault.NewMaster()

				// key_a, _ = vault.NewMaster()
				// key_b, _ = vault.NewMaster()

				// address_a, _ = vault.GetEncodedAddress(key_a)
				// address_b, _ = vault.GetEncodedAddress(key_b)
			})
		})

		Convey("GenerateIntoAddress()", func() {
			Convey("Should send bitcoins from a to b", func() {
				key_a, _ := vault.NewMaster()
				// TODO these should probably return a potential error.
				address_a := vault.GetEncodedAddress(key_a)
				err := GenerateInto(0.5, address_a)
				So(err, ShouldBeNil)
			})
		})
	})
}
