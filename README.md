<div align="center">
  <p>
    <img width="100" src="./img/dart.png">
  </p>
  <img src="https://img.shields.io/github/license/teodev1611/darth-db?style=flat-square">
  <img src="https://img.shields.io/github/stars/teodev1611/darth-db?style=social">
  <img src="https://img.shields.io/github/go-mod/go-version/teodev1611/darth-db/main?filename=go.mod">
  <h1>The Darth Database for the Dark Side</h1>
  <hr>
</div>

## 🤔 What is that?

This is a tiny database for small projects using many formats inspired in **lowdb**

## 💻 How install?

```
go get github.com/TeoDev1611/darth-db
```

## Examples 👌

Write a json database with the data:

```go
package main

import (
	"github.com/TeoDev1611/darth-db/json"
)

func main() {
	data := map[string]interface{}{
		"database":    "darth-db",
		"awesome":     true,
		"easy":        true,
		"description": "A little db for the dark side",
	}
	json.WriteDB("sampledb.json", "  ", false, data)
}
```

## 💁 Todo

- [x] Write Json Files
- [x] Parse String to Map String interface
- [x] Get all data from the json
- [x] Get a single value of the json
- [x] Support for toml 
- [ ] Encrypt function
- [x] Examples
- [ ] Support for Yaml

## ✅ Authors

Special thanks to [GolangUA](https://github.com/GolangUA/gopher-logos) and [Red Panda](http://panda-art.red/) for this beautiful illustration 🤟. 
Project made by @TeoDev1611 
