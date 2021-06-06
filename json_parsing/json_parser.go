package json_parser

import (
	"encoding/json"

	"github.com/TeoDev1611/darth-db/errors"
)

func StringToMap(text string) map[string]interface{} {
	byteText := []byte(text)
	var dat map[string]interface{}
	err := json.Unmarshal(byteText, &dat)
	errors.CheckErrors(err)
	return dat
}
