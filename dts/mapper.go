package dts

import (
	"io"
	"github.com/smbody/common-server/app"
	"github.com/smbody/common-server/bl/user"
)

func Write(v interface{}) ([]byte) {
	return app.Write(v)
}

func ToUser(reader io.Reader) (man *user.User) {
	man = user.New()
	app.Read(reader, man)
	return
}

