package maputils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInterfaceValue(t *testing.T) {
	Convey("Given a map[string]string", t, func() {
		m := map[string]string{
			"a": "a",
			"b": "b",
		}
		StringToInterface(m)
	})
}
