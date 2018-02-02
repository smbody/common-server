package greetings

import (
	"github.com/smbody/common-server/bl/user"
	"github.com/smbody/common-server/app"
	"github.com/smbody/common-server/errors"
)

const defaultTreatment = "World"

type greetings struct {
	defaultTreatment string
}

func New() *greetings {
	return &greetings{defaultTreatment}
}

func (g *greetings) SayHello(user *user.User) (*Hello) {
	if user.Name == "" {app.ThrowError(errors.UnknownUser)}
	return NewHello(user.Name)
}

func (g *greetings) SayHelloWorld() (*Hello) {
	return NewHello(g.defaultTreatment)
}

func (g *greetings) Ping() (*statistics) {
	return newStatistics()
}

