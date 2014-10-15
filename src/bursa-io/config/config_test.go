package middleware

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("Config Tests", t, func() {

		Convey("LoadConfig", func() {
			loadConfig()
			asset_path := viper.Get("asset_path")
			So(asset_path, ShouldNotBeNil)
			So(asset_path, ShouldEqual, "../../dist/")
		})
	})

}
