package tests

import (
	"bytes"
	"github.com/smbody/common-server/app"
	"github.com/smbody/common-server/dts"
	"github.com/smbody/common-server/tests/mocks"
	"testing"
	"github.com/smbody/common-server/container"
)



func TestHelloWorld(t *testing.T) {

	expected := mocks.ToArray(container.Greetings().SayHelloWorld())

	if s := dts.HelloWorld(nil); string(s) != string(expected) {
		t.Errorf("Expected %q but has %q", expected, string(s))
	}
}

func TestStatistics(t *testing.T) {

	expected := mocks.ToArray(container.Greetings().Ping())

	if s := dts.Ping(nil); string(s) != string(expected) {
		t.Errorf("Expected %q but has %q", expected, string(s))
	}
}

func TestHelloEmptyUser(t *testing.T) {
	emptyReader := bytes.NewReader([]byte{})
	mocks.AwaitPanic(t, func() { dts.Hello(emptyReader) }, app.CantDecodeJson)
}

func TestHelloUser(t *testing.T) {
	user:=container.User()
	expected := mocks.ToArray(container.Greetings().SayHello(user))
	reader := bytes.NewReader(app.Write(user))
	if s := dts.Hello(reader); string(s) != string(expected) {
		t.Errorf("Expected %q but has %q", expected, string(s))
	}
}
