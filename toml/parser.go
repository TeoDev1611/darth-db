package toml

import (
	"github.com/TeoDev1611/darth-db/errors"
	toml "github.com/pelletier/go-toml/v2"
)

func StringToMap(text []byte) map[string]interface{} {
	var data map[string]interface{}
	err := toml.Unmarshal(text, &data)
	errors.CheckErrors(err)
	return data
}
