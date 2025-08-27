package service

import (
    "context"
    "errors"
    "net/http"
    "os"
    "time"

    "task-service/internal/model"
    "task-service/internal/repository"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskService struct {
    repo          repository.TaskRepository
    userSvcBase   string
}

func NewTaskService(r repository.TaskRepository) *TaskService {
    base := os.Getenv("USER_SERVICE_BASE")
    if base == "" {
        base = "http://user-service:8081"
    }
    return &TaskService{repo: r, userSvcBase: base}
}

func (s *TaskService) Create(ctx context.Context, t *model.Task) error {
    now := time.Now()
    t.ID = primitive.NewObjectID()
    t.CreatedAt = now
    t.UpdatedAt = now

    if t.Status == "" {
        t.Status = model.StatusTodo
    }

    // validate userId if supplied
    if t.UserID != nil {
        exists := s.userExists(ctx, *t.UserID)
        if !exists {
            return errors.New("user_not_found")
        }
    }

    return s.repo.Create(ctx, t)
}

func (s *TaskService) Get(ctx context.Context, id primitive.ObjectID) (*model.Task, error) {
    return s.repo.GetByID(ctx, id)
}

func (s *TaskService) Update(ctx context.Context, t *model.Task) error {
    // if userId field present, validate
    if t.UserID != nil {
        if !s.userExists(ctx, *t.UserID) {
            return errors.New("user_not_found")
        }
    }
    t.UpdatedAt = time.Now()
    return s.repo.Update(ctx, t)
}

func (s *TaskService) Delete(ctx context.Context, id primitive.ObjectID) error {
    return s.repo.Delete(ctx, id)
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

// userExists checks user-service for the given ID
func (s *TaskService) userExists(ctx context.Context, id primitive.ObjectID) bool {
    url := s.userSvcBase + "/api/v1/users/" + id.Hex()
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
    if err != nil {
        return false
    }
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return false
    }
    resp.Body.Close()
    return resp.StatusCode == http.StatusOK
}
