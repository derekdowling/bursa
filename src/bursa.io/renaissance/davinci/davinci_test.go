package davinci

import (
	"bursa.io/renaissance/satchel"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// A little helper struct that allows you to throw a bunch of arbitrary anonymous
// functions into it for testing purposes
type FunctionRunner struct {
	Functions []func()
	stateful  bool
}

// Implements Handle so our Divinci Contraption can run it
func (self *FunctionRunner) Trigger(s *satchel.Satchel) {
	for _, function := range self.Functions {
		function()
	}
}

// Allows you to add functions more easily
func (self *FunctionRunner) Add(invokable func()) {
	self.Functions = append(self.Functions, invokable)
}

func TestSpec(t *testing.T) {

	Convey("Davinci Tests", t, func() {

		Convey("Function Runner", func() {
			runner := new(FunctionRunner)
			runner.Add(func() {})

			So(len(runner.Functions), ShouldEqual, 1)
		})

		blueprint := new(Blueprint)
		testMechanism := new(FunctionRunner)
		testController := new(FunctionRunner)

		Convey("AddMechanisms()", func() {
			Convey("Should store properly", func() {
				blueprint.AddMechanisms([]Mechanism{testMechanism})
				So(len(blueprint.mechanisms), ShouldEqual, 1)
			})
		})

		Convey("AddController()", func() {
			blueprint.AddController(testController)
			So(blueprint.controller, ShouldNotBeNil)

			// This test needs some magic to catch the Panic instead of erroring
			// out the test
			So(func() { blueprint.AddController(testController) }, ShouldPanic)
		})

		Convey("Test Fluency", func() {
			chained_blueprint := new(Blueprint)
			chained_blueprint_2 := chained_blueprint.
				AddMechanisms([]Mechanism{testMechanism}).
				AddController(testController)
			So(chained_blueprint, ShouldHaveSameTypeAs, blueprint)
			So(chained_blueprint, ShouldEqual, chained_blueprint_2)
		})

		Convey("build()", func() {
			blueprint := new(Blueprint)
			contraption := blueprint.
				AddMechanisms([]Mechanism{testMechanism}).
				AddController(testController).
				build()

			So(contraption, ShouldHaveSameTypeAs, &Contraption{})
			So(len(contraption.mechanisms), ShouldEqual, 2)
			So(contraption.mechanisms[0], ShouldNotEqual, testMechanism)
		})
	})
}
