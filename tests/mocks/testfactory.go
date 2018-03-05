package mocks

import (
	"github.com/smbody/common-server/bl"
)

type TestFactory struct {
}

func (f *TestFactory) Greetings() bl.Greetings {
	return &GreetingsMock{}
}

func (f *TestFactory) User() bl.User {
	return &UserMock{Name:"StarMan"}
}

