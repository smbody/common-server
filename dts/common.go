package dts

import (
	"github.com/smbody/common-server/app"
	"github.com/smbody/common-server/bl"
	"github.com/smbody/common-server/container"
	"io"
)

//Преобразование любой структуры в набор байтов.
func Write(v interface{}) []byte {
	return app.Write(v)
}

//Десереализация 
func ToUser(reader io.Reader) (man bl.User) {
	man = container.User()
	app.Read(reader, man)
	return
}
