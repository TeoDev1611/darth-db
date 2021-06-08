package toml

import (
	"github.com/TeoDev1611/darth-db/errors"
	toml "github.com/pelletier/go-toml"
)

func StringToMap(text []byte) map[string]interface{} {
	var data map[string]interface{}
	val, err := toml.LoadBytes(text)
	errors.CheckErrors(err)
	val.Unmarshal(&data)
	return data
}
