package reflectutils

import (
	"reflect"
)

func DeepValue(v reflect.Value) reflect.Value {
	k := v.Kind()
	switch k {
	case reflect.Ptr, reflect.Interface:
		el := v.Elem()
		return DeepValue(el)
	}
	return v
}
