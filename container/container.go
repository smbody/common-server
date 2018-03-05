package container

import "github.com/smbody/common-server/bl"

var (
	factory Factory

	greetingsBl bl.Greetings
)


func instance() Factory {
	if factory == nil {
		factory = &appFactory{}
	}
	return factory
}

func Init(f Factory) {
	factory = f
}

func Greetings() bl.Greetings {
	if greetingsBl == nil {
		greetingsBl = instance().Greetings()
	}
	return greetingsBl
}

func User() bl.User {
	return instance().User()
}
