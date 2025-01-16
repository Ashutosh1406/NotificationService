package utils

import "encoding/json"

func ToJson(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(data)
}
