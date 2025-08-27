package mongo

import (
    "context"
    "errors"

    "task-service/internal/model"
    repo "task-service/internal/repository"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type taskRepo struct {
    col *mongo.Collection
}

func NewTaskRepository(col *mongo.Collection) repo.TaskRepository {
    return &taskRepo{col: col}
}

func (r *taskRepo) Create(ctx context.Context, task *model.Task) error {
    _, err := r.col.InsertOne(ctx, task)
    return err
}

func (r *taskRepo) GetByID(ctx context.Context, id primitive.ObjectID) (*model.Task, error) {
    var t model.Task
    err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&t)
    if err == mongo.ErrNoDocuments {
        return nil, errors.New("task not found")
    }
    return &t, err
}

func (r *taskRepo) List(ctx context.Context, filter repo.TaskFilter, pg repo.Pagination) ([]*model.Task, int64, error) {
    query := bson.M{}
    if filter.Status != nil {
        query["status"] = *filter.Status
    }
    opts := options.Find().SetLimit(pg.Limit).SetSkip(pg.Offset)
    cur, err := r.col.Find(ctx, query, opts)
    if err != nil {
        return nil, 0, err
    }
    defer cur.Close(ctx)
    var tasks []*model.Task
    if err = cur.All(ctx, &tasks); err != nil {
        return nil, 0, err
    }
    total, _ := r.col.CountDocuments(ctx, query)
    return tasks, total, nil
}

func (r *taskRepo) Update(ctx context.Context, task *model.Task) error {
    _, err := r.col.ReplaceOne(ctx, bson.M{"_id": task.ID}, task)
    return err
}

func (r *taskRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
    _, err := r.col.DeleteOne(ctx, bson.M{"_id": id})
    return err
}
