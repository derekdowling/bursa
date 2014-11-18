package config

import (
	"github.com/derekdowling/mamba"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("Config Tests", t, func() {

		Convey("getLoadPath()", func() {

			path := getLoadPath()
			So(path, ShouldNotBeNil)
			So(path, ShouldContainSubstring, "config/yml")

			Convey("LoadDB()", func() {
				config := LoadDB(path)
				So(config, ShouldNotBeNil)
				So(config, ShouldHaveSameTypeAs, new(mamba.Config))
				So(config.GetStringMapString("orm")["adapter"], ShouldEqual, "postgres")
			})

			Convey("LoadServer()", func() {
				original_env := os.Getenv("BURSA_ENV")

				Convey("With Blank or Development Environment", func() {
					os.Setenv("BURSA_ENV", "")
					config := LoadServer(path)

					So(config, ShouldNotBeNil)
					So(config, ShouldHaveSameTypeAs, new(mamba.Config))
					asset_path := config.GetStringMapString("paths")["assets"]
					So(asset_path, ShouldNotBeNil)
					So(asset_path, ShouldEqual, "./assets")
				})

				Convey("With Production Environment", func() {
					os.Setenv("BURSA_ENV", "production")
					config := LoadServer(path)
					So(config.GetStringMap("logging")["mode"], ShouldEqual, "prod")
				})

				os.Setenv("BURSA_ENV", original_env)
			})

			Convey("LoadConfig()", func() {
				LoadConfig()
				So(Server, ShouldHaveSameTypeAs, new(mamba.Config))
				So(DB, ShouldHaveSameTypeAs, new(mamba.Config))

				// test reload prevention
				Server.Set("test123", "woo")
				LoadConfig()
				So(Server.Get("test123"), ShouldEqual, "woo")
			})
		})
	})

}
