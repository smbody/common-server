package errors

import (
	"github.com/smbody/common-server/app"
	"net/http"
)

const (
	UnknownUser app.ErrorCode = 1000
)

func init() {
	app.PredefineError(UnknownUser, http.StatusInternalServerError, "User must have a name. Bye!")
}
