package mocks

import (
	"github.com/smbody/common-server/bl"
)

type GreetingsMock struct {
}

func (g *GreetingsMock) SayHello(user bl.User) (bl.Hello) {
	return &HelloMock{Value: "hello test, "+user.FullName()+"!"}
}

func (g *GreetingsMock) SayHelloWorld() (bl.Hello) {
	return &HelloMock{Value: "hello test, World!"}
}

func (g *GreetingsMock) Ping() (bl.Statistics) {
	return NewStatisticsMock()
}

