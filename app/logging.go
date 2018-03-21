package app

import (
	"log"
	"net/http"
	"time"
)

func (app *application) logging(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%v] %q %v\n", r.Method, r.URL, t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}
