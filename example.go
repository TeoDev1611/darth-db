package main

import (
	"github.com/TeoDev1611/darth-db/files"
)

func main() {
	/*data := json_parser.StringToMap(`

	{"page": 1, "fruits": ["apple", "peach"]}
	`)*/
	/* 	var data = map[string]interface{} */
	files.WriteJsonDB("", "", true)
}
