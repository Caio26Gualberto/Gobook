package routes

import "net/http"

type Route struct {
	Uri                     string
	Method                  string
	Function                func(http.ResponseWriter, *http.Request)
	IsRequireAuthentication bool
}
