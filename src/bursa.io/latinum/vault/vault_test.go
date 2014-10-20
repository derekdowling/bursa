package vault

import (
	"bursa.io/models"
	"bursa.io/testutils"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"log"
)

func TestSpec(t *testing.T) {
	// TODO Doing this all over is ugly.
	db, err := models.Connect()
	if err != nil {
		log.Fatalf("Couldn't connect to database during testing", err)
	}

	Convey("Vault Tests", t, func() {
		Convey("NewMaster()", func() {
			Convey("Should generate a new HD Master Key.", func() {
				user := models.User{
					Name: testutils.SuffixedId("hd_key"),
				}
				So(db.Save(&user).Error, ShouldBeNil)

				new_master := NewMaster(user.Id)
				So(new_master, ShouldHaveSameTypeAs, "")
				// TODO assert key lenght is correct.
			})
		})
	})
}
