package assertions

import "github.com/gocql/gocql"

func ShouldBeUUID(actual interface{}, expected ...interface{}) string {
	uuid := actual.(string)
	_, err := gocql.ParseUUID(uuid)
	if err != nil {
		return "Expected an UUID"
	}
	return ""
}
