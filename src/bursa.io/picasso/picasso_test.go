package picasso

import (
	"bursa.io/config"
	. "github.com/smartystreets/goconvey/convey"
	"path"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("Picasso Tests", t, func() {

		Convey("getTemplateRoot", func() {
			temp := getTemplateRoot()
			So(temp, ShouldNotBeNil)
			So(path.Base(temp), ShouldEqual, "views")
		})

		Convey("findPartials()", func() {

			config.LoadConfig()
			template_path := getTemplateRoot()
			layout_path := "marketing/layout"

			Convey("should work with a good path", func() {

				layout := path.Join(template_path, layout_path)
				partials := findPartials(layout)

				So(partials, ShouldNotBeNil)
				So(len(partials), ShouldBeGreaterThan, "../../dist/")
			})
		})
	})
}
