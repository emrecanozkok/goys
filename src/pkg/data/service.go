package data

func GetValue(key string) (KeyValue, error) {
	var result KeyValue
	if _, ok := GlobalStore[key]; ok {
		result.Key = key
		result.Value = GlobalStore[key]
	}
	return result, nil
}

func SetValue(data KeyValue) bool {
	GlobalStore[data.Key] = data.Value
	return true
}
