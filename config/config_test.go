package config

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("Config Tests", t, func() {

		Convey("LoadConfig()", func() {
			LoadConfig()
			asset_path := viper.GetStringMapString("paths")["assets"]
			So(asset_path, ShouldNotBeNil)
			So(asset_path, ShouldEqual, "./assets")
		})

		Convey("GetString()", func() {
			asset_path := GetStringMapString("paths")["assets"]
			So(asset_path, ShouldNotBeNil)
			So(asset_path, ShouldEqual, "./assets")
		})
	})

}
