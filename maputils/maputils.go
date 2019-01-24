package maputils

func StringToInterface(stringMap map[string]string) map[string]interface{} {
	ifaceMap := map[string]interface{}{}
	for k, v := range stringMap {
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
