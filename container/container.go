//Package container - альтернатива контейнеру зависимостей.
//Все объекты, которые необходимы для работы приложения создаются и хранятся здесь.
//Здесь реализуется логика создания нового экземпляра класса или переиспользование существующего.
package container

import (
	"github.com/smbody/common-server/bl"
)

var (
	factory Factory

	greetingsBl bl.Greetings
)

//Возвращает экземпляр фабрики, реализующей создание объектов
func instance() Factory {
	if factory == nil {
		factory = &appFactory{}
	}
	return factory
}

//Init - используется для подмены фабрики, отвечающей за создание новых экземпляров объектов.
//Например, в тестах он вызывается с экземпляром фабрики, которая возвращет mock'и для всех классов бизнес-логики.
func Init(f Factory) {
	factory = f
}

//Greetings - инициализация экземпляра класса, отвечающего за приветственные сообщения сервера.
func Greetings() bl.Greetings {
	if greetingsBl == nil {
		greetingsBl = instance().Greetings()
	}
	return greetingsBl
}

//User - получение экземпляра класса бизнес-логики, отвечающего за работу с пользователем
func User() bl.User {
	return instance().User()
}
