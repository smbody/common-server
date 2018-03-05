package greetings

import (
	"github.com/smbody/common-server/app"
	"github.com/smbody/common-server/errors"
	"github.com/smbody/common-server/bl"
)

const defaultTreatment = "World"

type greetings struct {
	defaultTreatment string
}

func New() *greetings {
	return &greetings{defaultTreatment}
}

func (g *greetings) SayHello(user bl.User) (bl.Hello) {
	if user.FullName() == "" {app.ThrowError(errors.UnknownUser)}
	return NewHello(user.FullName())
}

func (g *greetings) SayHelloWorld() (bl.Hello) {
	return NewHello(g.defaultTreatment)
}

func (g *greetings) Ping() (bl.Statistics) {
	return newStatistics()
}

