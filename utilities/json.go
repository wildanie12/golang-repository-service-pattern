package utilities

import "encoding/json"

func JsonEncode(data interface{}) string {
	json, _ := json.MarshalIndent(data, "", "	")
	return string(json)
}