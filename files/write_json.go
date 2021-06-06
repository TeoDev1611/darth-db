package files

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"github.com/TeoDev1611/darth-db/errors"
)

func WriteJsonDB(name, spaces string, empty bool, data ...map[string]interface{}) {
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

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func EncryptFile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, passphrase))
}
