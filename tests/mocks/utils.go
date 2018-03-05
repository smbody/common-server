package mocks

import (
	"encoding/json"
	"log"
)

func ToArray(v interface{})  []byte {
	if result, err := json.Marshal(v); err==nil {
		return result
	}
	log.Fatalf("Can't marshal object: %v", v)
	return []byte{}
}


