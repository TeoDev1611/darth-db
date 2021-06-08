package json

import (
	"encoding/json"

	"github.com/TeoDev1611/darth-db/errors"
)

func StringToMap(text string) map[string]interface{} {
	byteText := []byte(text)
	var data map[string]interface{}
	err := json.Unmarshal(byteText, &data)
	errors.CheckErrors(err)
	return data
}
