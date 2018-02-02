package app

import (
	"net/http"
	"fmt"
)

type ErrorCode int

type serverError struct {
	errCode ErrorCode
	status  int
	errText string
}

var predefinedErrors map[ErrorCode]*serverError

const (
	UnknownError      ErrorCode = 1
	CantDecodeJson              = 2
	CantMarshalObject           = 3
)

func init() {
	predefinedErrors = make(map[ErrorCode]*serverError)
	PredefineError(UnknownError, http.StatusInternalServerError, "Internal server error")
	PredefineError(CantDecodeJson, http.StatusBadRequest, "Can't decode json")
	PredefineError(CantMarshalObject, http.StatusInternalServerError, "Can't marshal object")
}

func PredefineError(errCode ErrorCode, status int, errText string) {
	predefinedErrors[errCode] = &serverError{errCode, status, errText}
}

func ThrowError(errCode ErrorCode) {
	panic(predefinedErrors[errCode])
}

func (e *serverError) Error() string {
	return fmt.Sprintf("%v: %s", e.errCode, e.errText)
}
