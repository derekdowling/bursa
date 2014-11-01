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

		Convey("getTemplateRoot()", func() {
			template_root := getTemplateRoot()
			So(template_root, ShouldNotBeNil)
			So(path.Base(template_root), ShouldEqual, "views")

			files, _ := filepath.Glob(template_root + "/*")
			So(len(files), ShouldBeGreaterThan, 0)
		})

		// Used By findPartials() and parseTemplates()
		config.LoadConfig()
		template_path := getTemplateRoot()
		layout := "marketing/layout.tmpl"
		layout_path := path.Join(template_path, layout)
		view_path := path.Join(template_path, "marketing/index.tmpl")

		Convey("findPartials()", func() {

			Convey("should work with a good path", func() {

				partials := findPartials(layout_path)

				partial_count := len(partials)

				So(partials, ShouldNotBeNil)
				So(partial_count, ShouldBeGreaterThan, 0)

				// ensure that signup is actually a partial
				signup_index := sort.Search(partial_count, func(i int) bool { return strings.HasSuffix(partials[i], "signup.tmpl") })
				So(signup_index, ShouldNotEqual, partial_count)
			})

			Convey("should not explode on an empty route", func() {
				layout := path.Join(template_path, "notreal/layout")
				partials := findPartials(layout)

				So(len(partials), ShouldEqual, 0)
			})
		})

		Convey("combineTemplates()", func() {

			Convey("should successfully compile a template glob", func() {
				template := combineTemplates(layout_path, view_path, findPartials(layout_path))

				So(template, ShouldNotBeNil)
				So(template, ShouldHaveSameTypeAs, template.New("example"))
				So(template.Lookup("content"), ShouldNotBeNil)
				So(template.Lookup("signup"), ShouldNotBeNil)
			})

			Convey("should successfully compile if no partials are present", func() {
				template := combineTemplates(layout_path, view_path, findPartials("no_partials_path"))

				So(template, ShouldNotBeNil)
				So(template, ShouldHaveSameTypeAs, template.New("example"))
				So(template.Lookup("content"), ShouldNotBeNil)
				So(template.Lookup("signup"), ShouldBeNil)
			})
		})

		Convey("buildTemplate()", func() {

			Convey("should successfully build a template", func() {
				template := buildTemplate("marketing/layout", "marketing/index")

				So(template, ShouldNotBeNil)
				So(template, ShouldHaveSameTypeAs, template.New("example"))
				So(template.Lookup("content"), ShouldNotBeNil)
				So(template.Lookup("signup"), ShouldNotBeNil)
			})

		})
	})
}
