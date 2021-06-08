package files

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/TeoDev1611/darth-db/errors"
)

func WriteJSONDB(name, spaces string, empty bool, data ...map[string]interface{}) {
	if spaces == "" {
		spaces = "  "
	}

	if name == "" {
		name = "./db.json"
	}
	if empty {
		data = []map[string]interface{}{map[string]interface{}{
			"database":    "Darth DB",
			"awesome":     true,
			"description": "The best database of the world",
		}}
	}
	file, err := json.MarshalIndent(data, "", spaces)
	errors.CheckErrors(err)
	err2 := ioutil.WriteFile(name, file, os.ModePerm)
	errors.CheckErrors(err2)
}


