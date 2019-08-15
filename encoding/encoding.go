//This package provides functions to make easier encode and decode
//elements. It can be protos to structs and viceversa and more formats
package encoding

import "encoding/json"

//EncodeToProto parse a proto to the given struct. To Use this function is
//Important that the json tags in the proto and inside the struct must match.
//Also you have to pass out parameter as pointer of your struct
func DecodeFromProto(pb interface{}, out interface{}) error {
	var err error
	var data []byte

	if data, err = json.Marshal(pb); err != nil {
		return err
	}

	if err := json.Unmarshal(data, out); err != nil {
		return err
	}

	return nil
}

//EncodeToProto encode the given struct in the given proto (pb).
//pb parameter must be a pointer of proto
//Json tags between the given struct (model) and the proto (pb) must be equal.
func EncodeToProto(model interface{}, pb interface{}) error {
	return DecodeFromProto(model, pb)
}
