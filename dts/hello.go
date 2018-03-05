package dts

import (
	"io"
	"github.com/smbody/common-server/container"
)


func Hello(r io.Reader) (result []byte) {
	man := ToUser(r)
	hello := container.Greetings().SayHello(man)
	return Write(hello)
}

func HelloWorld(r io.Reader) ([]byte) {
	return Write(container.Greetings().SayHelloWorld())
}

func Ping(r io.Reader) (result []byte) {
	return Write(container.Greetings().Ping())
}
