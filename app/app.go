package app

import (
	"fmt"
	"github.com/justinas/alice"
	"log"
	"net/http"
)

var App *application

type application struct {
	routes         Routes
	defaultHandler http.Handler
	host           string
	port           string
}

func Init(routes Routes) *application {
	result := &application{
		routes: routes,
		host:   "",
		port:   "8001",
	}

	result.defaultHandler = alice.New(
		result.logging,
		result.methodAllowed,
	).Then(result.executor())

	return result
}

func Run() {
	for url := range App.routes {
		http.Handle(url, App.defaultHandler)
	}
	listen := fmt.Sprintf("%s:%s", App.host, App.port)
	log.Println("Server started (port:" + App.port+")")
	if err := http.ListenAndServe(listen, nil); err != nil {
		log.Println(err)
	}
}
