package routes

import (
	"net/http"
)

var routes []route

type route struct {
	Path    string
	Method  string
	Name    string
	Handler http.HandlerFunc
}
type group struct {
	prefix string
	routes []route
}
