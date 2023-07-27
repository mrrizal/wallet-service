package utils

import (
	"encoding/json"
)

func MapToJson(data map[string]interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func JsonToMap(jsonData string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return make(map[string]interface{}), err
	}
	return data, nil
}
