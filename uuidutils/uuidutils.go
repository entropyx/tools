package uuidutils

import "github.com/gocql/gocql"

func IsUUID(uuid string) bool {
	_, err := gocql.ParseUUID(uuid)
	if err != nil {
		return false
	}
	return true
}
