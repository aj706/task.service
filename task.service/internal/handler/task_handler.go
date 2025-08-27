package handler

import (
    "encoding/json"
    "net/http"
    "strconv"

    "task-service/internal/model"
    "task-service/internal/repository"
    "task-service/internal/service"

    "github.com/go-chi/chi/v5"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskHandler struct {
    svc *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
    return &TaskHandler{svc: s}
}

// Register wires task endpoints into provided router
func (h *TaskHandler) Register(r chi.Router) {
    r.Post("/tasks", h.create)
    r.Get("/tasks", h.list)
    r.Route("/tasks/{id}", func(sr chi.Router) {
        sr.Get("/", h.getOne)
        sr.Patch("/", h.update)
        sr.Delete("/", h.delete)
    })
}

func (h *TaskHandler) create(w http.ResponseWriter, r *http.Request) {
    var payload model.Task
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := h.svc.Create(r.Context(), &payload); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(payload)
}

func (h *TaskHandler) list(w http.ResponseWriter, r *http.Request) {
    limitStr := r.URL.Query().Get("limit")
    offsetStr := r.URL.Query().Get("offset")
    statusStr := r.URL.Query().Get("status")

    var limit int64 = 20
    var offset int64 = 0
    if l, err := strconv.ParseInt(limitStr, 10, 64); err == nil {
        limit = l
    }
    if o, err := strconv.ParseInt(offsetStr, 10, 64); err == nil {
        offset = o
    }

    var filter repository.TaskFilter
    if statusStr != "" {
        st := model.TaskStatus(statusStr)
        filter.Status = &st
    }

    tasks, total, err := h.svc.List(r.Context(), filter, repository.Pagination{Limit: limit, Offset: offset})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    resp := map[string]interface{}{
        "items":  tasks,
        "total":  total,
        "limit":  limit,
        "offset": offset,
    }
    json.NewEncoder(w).Encode(resp)
}

func (h *TaskHandler) getOne(w http.ResponseWriter, r *http.Request) {
    idHex := chi.URLParam(r, "id")
    id, err := primitive.ObjectIDFromHex(idHex)
    if err != nil { http.Error(w, "invalid id", 400); return }
    task, err := h.svc.Get(r.Context(), id)
    if err != nil { http.Error(w, err.Error(), 404); return }
    json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) update(w http.ResponseWriter, r *http.Request) {
    idHex := chi.URLParam(r, "id")
    id, err := primitive.ObjectIDFromHex(idHex)
    if err != nil { http.Error(w, "invalid id", 400); return }
    var payload model.Task
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, err.Error(), 400); return }
    payload.ID = id
    if err := h.svc.Update(r.Context(), &payload); err != nil {
        if err.Error() == "user_not_found" {
            http.Error(w, "user not found", 400)
        } else {
            http.Error(w, err.Error(), 500)
        }
        return
    }
    json.NewEncoder(w).Encode(payload)
}

func (h *TaskHandler) delete(w http.ResponseWriter, r *http.Request) {
    idHex := chi.URLParam(r, "id")
    id, err := primitive.ObjectIDFromHex(idHex)
    if err != nil { http.Error(w, "invalid id", 400); return }
    if err := h.svc.Delete(r.Context(), id); err != nil {
        http.Error(w, err.Error(), 500); return }
    w.WriteHeader(204)
}