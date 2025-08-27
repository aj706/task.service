package service

import (
    "context"
    "time"

    "task-service/internal/model"
    "task-service/internal/repository"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskService struct {
    repo repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) *TaskService {
    return &TaskService{repo: r}
}

func (s *TaskService) Create(ctx context.Context, t *model.Task) error {
    now := time.Now()
    t.ID = primitive.NewObjectID()
    t.CreatedAt = now
    t.UpdatedAt = now
    if t.Status == "" {
        t.Status = model.StatusPending
    }
    return s.repo.Create(ctx, t)
}

func (s *TaskService) List(ctx context.Context, filter repository.TaskFilter, pg repository.Pagination) ([]*model.Task, int64, error) {
    if pg.Limit <= 0 || pg.Limit > 100 {
        pg.Limit = 20
    }
    if pg.Offset < 0 {
        pg.Offset = 0
    }
    return s.repo.List(ctx, filter, pg)
}
