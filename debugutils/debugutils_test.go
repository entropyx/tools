package debugutils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStack(t *testing.T) {
	Convey("When a stack is generated", t, func() {
		s := Stack()
		Convey("The stack should include this test", func() {
			So(s, ShouldContainSubstring, "debugutils_test")
		})
	})
}

func TestStackSimple(t *testing.T) {
	Convey("When a stack is generated", t, func() {
		s := StackSimple()
		Convey("The stack should include this test", func() {
			So(s, ShouldContainSubstring, "debugutils_test")
		})

		Convey("The stack should not include the stack functions", func() {
			So(s, ShouldNotContainSubstring, stackFuncRuntime)
			So(s, ShouldNotContainSubstring, stackFuncDebugutils)
		})
	})
}
