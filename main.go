package main

import (
	"github.com/smbody/common-server/application"
	"github.com/smbody/common-server/services"
)

func main() {
	Application := application.Init(application.Routes{
		"/calculate": services.Calculate ,
		"/graphql": application.NewGraphqQL(services.GraphQLSchema),
		})
	Application.Run()
}
