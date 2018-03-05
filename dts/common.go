package dts

import (
	"github.com/smbody/common-server/app"
	"github.com/smbody/common-server/container"
	"github.com/smbody/common-server/bl"
	"io"
)

func Write(v interface{}) []byte {
	return app.Write(v)
}

func ToUser(reader io.Reader) (man bl.User) {
	man = container.User()
	app.Read(reader, man)
	return
}
