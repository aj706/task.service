package repository

import (
    "context"

    "task-service/internal/model"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Pagination struct {
    Limit  int64
    Offset int64
}

type TaskFilter struct {
    Status *model.TaskStatus
}

type TaskRepository interface {
    Create(ctx context.Context, task *model.Task) error
    GetByID(ctx context.Context, id primitive.ObjectID) (*model.Task, error)
    List(ctx context.Context, filter TaskFilter, pg Pagination) ([]*model.Task, int64, error)
    Update(ctx context.Context, task *model.Task) error
    Delete(ctx context.Context, id primitive.ObjectID) error
}
