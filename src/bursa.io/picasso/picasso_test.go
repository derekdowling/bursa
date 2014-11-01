package picasso

import (
	"bursa.io/config"
	. "github.com/smartystreets/goconvey/convey"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("Picasso Tests", t, func() {

		Convey("getTemplateRoot", func() {
			template_root := getTemplateRoot()
			So(template_root, ShouldNotBeNil)
			So(path.Base(template_root), ShouldEqual, "views")

			files, _ := filepath.Glob(template_root + "/*")
			So(len(files), ShouldBeGreaterThan, 0)
		})

		Convey("findPartials()", func() {

			config.LoadConfig()
			template_path := getTemplateRoot()
			layout_path := "marketing/layout"

			Convey("should work with a good path", func() {

				layout := path.Join(template_path, layout_path)
				partials := findPartials(layout)

				partial_count := len(partials)

				So(partials, ShouldNotBeNil)
				So(partial_count, ShouldBeGreaterThan, 0)

				// ensure that signup is actually a partial
				signup_index := sort.Search(partial_count, func(i int) bool { return strings.HasSuffix(partials[i], "signup.tmpl") })
				So(signup_index, ShouldNotEqual, partial_count)
			})

			Convey("not explode on an empty route", func() {
				layout := path.Join(template_path, "notreal/layout")
				partials := findPartials(layout)

				So(len(partials), ShouldEqual, 0)
			})
		})
	})
}
