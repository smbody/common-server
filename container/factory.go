package container

import (
	"github.com/smbody/common-server/bl"
)

//Factory - интерфейс, определяющий требования к структуре, реализующей фабрику инициализации объектов приложения.
type Factory interface {
	Greetings() bl.Greetings
	User() bl.User
}
