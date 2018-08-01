package maputils

func StringToInterface(stringMap map[string]string) map[string]interface{} {
	ifaceMap := map[string]interface{}{}
	for k, v := range stringMap {
		ifaceMap[k] = v
	}
	return ifaceMap
}
