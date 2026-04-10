package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router chi.Router
}

func NewServer() *Server {
	s := &Server{}
	s.createRouter()
	return s
}

func (s *Server) createRouter() {
	s.Router = chi.NewRouter()

	s.Router.Route("/auth/", func(r chi.Router) {
		r.Post("/register", RegisterUser)
		//r.Post("/login")
	})
}

func (s *Server) ListenAndServe() error {
	return http.ListenAndServe("", s.Router)
}
