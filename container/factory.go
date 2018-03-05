package container

import (
	"github.com/smbody/common-server/bl"
)

type Factory interface {
	Greetings() bl.Greetings
	User() bl.User
}
