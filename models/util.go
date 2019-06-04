package models

import (
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

type Expirable interface {
	CreatedAt() int64
	ExpiresIn() int
}

type Cache interface {
	Set(data Expirable) error
	Get(data Expirable) error
}

type FileCache struct {
	Path string
}

func NewFileCache(path string) *FileCache {
	return &FileCache{
		Path: path,
	}
}

func (f *FileCache) Set(data Expirable) error {
	bytes, err := json.Marshal(data)
	if err == nil {
		ioutil.WriteFile(f.Path, bytes, 0644)
	}
	return err
}

func (f *FileCache) Get(data Expirable) error {
	bytes, err := ioutil.ReadFile(f.Path)
	if err == nil {
		err = json.Unmarshal(bytes, data)
		if err == nil {
			created := data.CreatedAt()
			expires := data.ExpiresIn()
			if err == nil && time.Now().Unix() > created+int64(expires-60) {
				err = errors.New("Data is already expired")
			}
		}
	}
	return err
}

func sha1Sign(s string) string {
	// The pattern for generating a hash is `sha1.New()`,
	// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
	// Here we start with a new hash.
	h := sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	h.Write([]byte(s))

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	bs := h.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	return fmt.Sprintf("%x", bs)
}

func HandJSONTopResponse(responseData interface{}, content []byte) {
	json.Unmarshal(content, responseData)
}

func HandXMLTopResponse(responseData interface{}, content []byte) {
	xml.Unmarshal(content, responseData)
}
