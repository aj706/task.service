package handler

import (
    "encoding/json"
    "net/http"
    "user-service/internal/model"
    "user-service/internal/service"
    "github.com/go-chi/chi/v5"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct { svc *service.UserService }
func New(s *service.UserService) *UserHandler { return &UserHandler{svc: s} }

func (h *UserHandler) Register(r chi.Router) {
    r.Post("/users", h.create)
    r.Route("/users/{id}", func(sr chi.Router) {
        sr.Get("/", h.get)
        sr.Put("/", h.update)
        sr.Delete("/", h.delete)
    })
}

func (h *UserHandler) create(w http.ResponseWriter, r *http.Request) {
    var u model.User
    if err := json.NewDecoder(r.Body).Decode(&u); err != nil { http.Error(w, err.Error(), 400); return }
    if err := h.svc.Create(r.Context(), &u); err != nil { http.Error(w, err.Error(), 500); return }
    w.WriteHeader(201); json.NewEncoder(w).Encode(u)
}

func (h *UserHandler) get(w http.ResponseWriter, r *http.Request) {
    id, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
    u, err := h.svc.Get(r.Context(), id)
    if err != nil { http.Error(w, err.Error(), 404); return }
    json.NewEncoder(w).Encode(u)
}

func (h *UserHandler) update(w http.ResponseWriter, r *http.Request) {
    id, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
    var u model.User; if err := json.NewDecoder(r.Body).Decode(&u); err != nil { http.Error(w, err.Error(),400); return }
    u.ID = id
    if err := h.svc.Update(r.Context(), &u); err != nil { http.Error(w, err.Error(),500); return }
    json.NewEncoder(w).Encode(u)
}

func (h *UserHandler) delete(w http.ResponseWriter, r *http.Request) {
    id, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
    if err := h.svc.Delete(r.Context(), id); err != nil { http.Error(w, err.Error(),500); return }
    w.WriteHeader(204)
}
