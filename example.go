package main

import (
	"fmt"

	"github.com/TeoDev1611/darth-db/data"
)

func main() {
	/*data := json_parser.StringToMap(`

	{"page": 1, "fruits": ["apple", "peach"]}
	`)*/
	/* 	var data = map[string]interface{} */

	dataMap := data.GetAllDataFile("./db.json")
	fmt.Print(dataMap)

	//	files.WriteJsonDB("", "  ", true)
}
