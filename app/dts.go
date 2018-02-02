package app

import (
	"encoding/json"
	"io"
)

func Read(reader io.Reader, entity interface{}) {
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(entity); err != nil {
		ThrowError(CantDecodeJson)
	}
}

func Write(v interface{}) []byte {
	if result, err := json.Marshal(v); err == nil {
		return result
	}
	ThrowError(CantMarshalObject)
	return nil
}
