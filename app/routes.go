package app

import (
	"net/http"
	"io"
)

type dtsMethod func(reader io.Reader) ([]byte)

type route struct {
	httpMethod string
	executor   dtsMethod
}

type Routes map[string] *route

func Post(executor dtsMethod) *route {
	return &route{http.MethodPost, executor}
}

func Get(executor dtsMethod) *route {
	return &route{http.MethodGet, executor}
}

func Any(executor dtsMethod) *route {
	return &route{executor:executor}
}

func (r *route) MethodAllowed(method string) bool {
	if r.httpMethod == "" || r.httpMethod == method {return true}
	return false
}
