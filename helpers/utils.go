package helpers

import (
	"bytes"
	"encoding/gob"
)

//GetRealSizeOf returns the number of bytes occupied by a given interface
func GetRealSizeOf(v interface{}) int {
	b := new(bytes.Buffer)
	err := gob.NewEncoder(b).Encode(v)
	if err != nil {
		panic(err)
	}
	return b.Len()
}
