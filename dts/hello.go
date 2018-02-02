package dts

import (
	"github.com/smbody/common-server/bl/greetings"
	"io"
)

var bl = greetings.New()

func Hello(r io.Reader) (result []byte) {
	man := ToUser(r)
	hello := bl.SayHello(man)
	return Write(hello)
}

func HelloWorld(r io.Reader) ([]byte) {
	return Write(bl.SayHelloWorld())
}

func Ping(r io.Reader) (result []byte) {
	return Write(bl.Ping())
}
