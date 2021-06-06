package data

import (
	"encoding/json"
	"os"

	"github.com/TeoDev1611/darth-db/errors"
)

func GetAllDataFile(filename string) []interface{} {
	file, err := os.ReadFile(filename)
	errors.CheckErrors(err)
	var jsonData []interface{}
	err2 := json.Unmarshal(file, &jsonData)
	errors.CheckErrors(err2)
	return jsonData
}

func GetSingleKey(filename, key string, position int) (string, interface{}) {
	jsonData := GetAllDataFile(filename)
	listInfo := jsonData[position]
	v, ok := listInfo.(map[string]interface{})
	if !ok {
		errors.NewError("No can transformate from interface{} to map[string]interface{}", true)
	}
	for k, val := range v {
		if k == key {
			return k, val
			break
		}
	}
	return "No data", false
}
