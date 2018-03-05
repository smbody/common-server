package container

import (
	"github.com/smbody/common-server/bl/greetings"
	"github.com/smbody/common-server/bl"
	"github.com/smbody/common-server/bl/user"
)

type appFactory struct {
}

func (f *appFactory) Greetings() bl.Greetings {
	return greetings.New()
}

func (f *appFactory) User() bl.User {
	return user.New()
}
