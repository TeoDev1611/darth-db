package files

import (
	"encoding/json"
	"io/ioutil"

	"github.com/TeoDev1611/darth-db/errors"
)

func WriteJsonDB(name string, data map[string]interface{}) {
	name = "db.json"
	data = map[string]interface{}{
		"database":    "Darth DB",
		"awesome":     true,
		"description": "The best database of the world",
	}
	file, err := json.MarshalIndent(data, "", "")
	errors.CheckErrors(err)
	err2 := ioutil.WriteFile(name, file, 0644)
	errors.CheckErrors(err2)
}
