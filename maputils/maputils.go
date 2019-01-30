package maputils

import (
	"strconv"
)

func StringToInterface(stringMap map[string]string) map[string]interface{} {
	ifaceMap := map[string]interface{}{}
	for k, v := range stringMap {
		if fv, err := strconv.ParseFloat(v, 64); err == nil {
			ifaceMap[k] = fv
			continue
		}
		ifaceMap[k] = v
	}
	return ifaceMap
}

func ToString(m map[string]string) (strMap string) {

	for k, v := range m {
		strMap += k + ":" + v + "."
	}

	if len(strMap) > 0 {
		strMap = strMap[:len(strMap)-1]
	}
	return
}
