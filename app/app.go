package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/justinas/alice"
)

var App *application

type application struct {
	routes         Routes
	auth           Auth
	defaultHandler http.Handler
	host           string
	port           string
}

func Init(auth Auth, routes Routes) *application {
	result := &application{
		routes: routes,
		auth:   auth,
		host:   "",
		port:   "8001",
	}

	result.defaultHandler = alice.New(
		result.logging,
		result.methodAllowed,
		//result.authenticator,
	).Then(result.executor())

	return result
}

func Run() {
	for url := range App.routes {
		http.Handle(url, App.defaultHandler)
	}
	listen := fmt.Sprintf("%s:%s", App.host, App.port)
	log.Println("Server started (port:" + App.port + ")")
	if err := http.ListenAndServe(listen, nil); err != nil {
		log.Println(err)
	}
}
