package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type option func(*app)
type app struct {
	*chi.Mux
	server *http.Server
}

func New(opts ...option) *app {
	a := &app{
		server: &http.Server{
			Addr:    ":3000",
			Handler: nil,
		},
		Mux: chi.NewRouter(),
	}
	for _, opt := range opts {
		opt(a)
	}
	a.boostrap()
	return a
}

func (a *app) Start() error {
	a.server.Handler = a.Mux
	return a.server.ListenAndServe()
}

func (a *app) WithRouter(f func(r *chi.Mux)) {
	f(a.Mux)
}
