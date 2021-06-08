package toml

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/TeoDev1611/darth-db/errors"
	"github.com/pelletier/go-toml/v2"
)

func WriteDB(name string, empty bool, data ...map[string]interface{}) {
	if name == "" {
		name = "db.toml"
	}
	if empty {
		data = []map[string]interface{}{map[string]interface{}{
			"database":    "Darth DB",
			"awesome":     true,
			"description": "The best database of the world",
		}}
	}
	buf := bytes.Buffer{}
	enc := toml.NewEncoder(&buf)
	enc.SetIndentTables(true)
	enc.Encode(data)
	err2 := ioutil.WriteFile(name, buf.Bytes(), os.ModePerm)
	errors.CheckErrors(err2)
}
