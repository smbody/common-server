package app

import (
	"log"
	"net/http"
)

func (app *application) executor() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		defer errorHandling(w)
		route := app.routes[req.URL.Path]
		result := route.executor(req.Body)
		w.Header().Set("Content-Type", "app/json")
		w.Write(result)
	}
	return http.HandlerFunc(fn)
}

func errorHandling(w http.ResponseWriter) {
	if e := recover(); e != nil {
		switch err := e.(type) {
		case *serverError:
			http.Error(w, err.Error(), err.status)
		case error:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("(errorHandling) Unhandled exception: %s", err.Error())
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Printf("(errorHandling) Waiting error but give: %s", err)
		}
	}

}
