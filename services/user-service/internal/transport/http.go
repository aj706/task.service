package transport

import (
    "log"
    "net/http"
    "user-service/internal/handler"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

type Server struct { router *chi.Mux }
func New(h *handler.UserHandler) *Server {
    r := chi.NewRouter(); r.Use(middleware.Logger, middleware.Recoverer)
    r.Route("/api/v1", func(rt chi.Router) { h.Register(rt) })
    return &Server{router: r}
}
func (s *Server) Start(addr string) { log.Fatal(http.ListenAndServe(addr, s.router)) }
