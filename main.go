package main

import (
	"github.com/smbody/common-server/app"
	"github.com/smbody/common-server/dts"
)

func main() {
	app.App = app.Init(app.Routes{
		"/hello": app.Get(dts.HelloWorld),
		"/user": app.Post(dts.Hello),
		"/ping": app.Any(dts.Ping),
		})
	app.Run()
}
