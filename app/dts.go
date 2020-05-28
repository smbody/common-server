package app

import (
	"encoding/json"
	"io"
)

//Read - Десереализация, восстановление начального состояния структуры данных из битовой последовательности.
func Read(reader io.Reader, entity interface{}) {
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(entity); err != nil {
		ThrowError(CantDecodeJson)
	}
}

//Write - Сериализация, конвертация структуры данных в последовательность битов.
func Write(v interface{}) []byte {
	if result, err := json.Marshal(v); err == nil {
		return result
	}
	ThrowError(CantMarshalObject)
	return nil
}
