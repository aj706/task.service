package transport

import (
    "log"
    "net/http"

    "task-service/internal/handler"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

type Server struct {
    router *chi.Mux
}

func NewServer(th *handler.TaskHandler) *Server {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    // Health check endpoint
    r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
    })

    r.Route("/api/v1", func(api chi.Router) {
        th.Register(api)
    })

    return &Server{router: r}
}

func (s *Server) Start(addr string) {
    log.Printf("listening on %s", addr)
    log.Fatal(http.ListenAndServe(addr, s.router))
}
