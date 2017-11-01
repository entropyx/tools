package assertions

import (
	"fmt"

	"github.com/entropyx/tools/uuidutils"
)

func ShouldBeUUID(actual interface{}, expected ...interface{}) string {
	uuid := actual.(string)
	if isUUID := uuidutils.IsUUID(uuid); isUUID == false {
		return fmt.Sprintf(shouldBeUUID, uuid)
	}
	return ""
}

func ShouldNotBeUUID(actual interface{}, expected ...interface{}) string {
	uuid := actual.(string)
	if isUUID := uuidutils.IsUUID(uuid); isUUID == true {
		return fmt.Sprintf(shouldNotBeUUID, uuid)
	}
	return ""
}
