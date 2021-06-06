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
